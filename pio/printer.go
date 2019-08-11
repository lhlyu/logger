package pio

import (
	"bufio"
	"bytes"
	"github.com/lhlyu/logger/pio/terminal"
	"io"
	"io/ioutil"
	"strconv"
	"sync"
	"sync/atomic"
)

type (
	Handler func(PrintResult)
)

type Printer struct {
	Name       string
	IsTerminal bool
	priority   int
	Chained    bool
	Output     io.Writer
	mu         sync.Mutex
	marshal    MarshalerFunc
	hijack     Hijacker
	handlers   []Handler
	io.Reader
	io.Writer
	io.Closer
	DirectOutput bool
}

var (
	TotalPrinters int32
)

func NewPrinter(name string, output io.Writer) *Printer {
	if output == nil {
		output = NopOutput()
	}
	atomic.AddInt32(&TotalPrinters, 1)

	if name == "" {
		totalPrinters := atomic.LoadInt32(&TotalPrinters)
		lens := strconv.Itoa(int(totalPrinters))
		name = "printer_" + lens
	}

	buf := &bytes.Buffer{}

	isOuputTerminal := isTerminal(output)

	p := &Printer{
		Name:       name,
		Output:     output,
		Writer:     buf,
		Reader:     buf,
		Closer:     NopCloser(),
		IsTerminal: isOuputTerminal,
	}
	return p
}

func NewTextPrinter(name string, output io.Writer) *Printer {
	p := NewPrinter(name, output)
	p.Marshal(Text)
	return p
}

func (p *Printer) Priority(prio int) *Printer {
	p.mu.Lock()
	p.priority = prio
	p.mu.Unlock()
	return p
}

func (p *Printer) Marshal(marshaler Marshaler) *Printer {
	return p.MarshalFunc(marshaler.Marshal)
}

func (p *Printer) MarshalFunc(marshaler func(v interface{}) ([]byte, error)) *Printer {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.marshal == nil {
		p.marshal = marshaler
		return p
	}

	oldM := p.marshal
	newM := marshaler

	p.marshal = func(v interface{}) ([]byte, error) {
		b, err := oldM(v)

		if err != nil && err.Error() == ErrMarshalNotResponsible.Error() {
			b, err = newM(v)
		}

		if len(b) <= 0 && err == nil {
			return b, ErrSkipped
		}
		return b, err
	}
	return p
}

func (p *Printer) WithMarshalers(marshalers ...Marshaler) *Printer {
	if len(marshalers) == 0 {
		return p
	}
	for _, marshaler := range marshalers {
		p.Marshal(marshaler)
	}
	return p
}

func (p *Printer) AddOutput(writers ...io.Writer) *Printer {
	p.mu.Lock()
	defer p.mu.Unlock()

	for _, w := range writers {
		if !terminal.IsTerminal(w) {
			p.IsTerminal = false
			break
		}
	}

	w := io.MultiWriter(append(writers, p.Output)...)
	p.Output = w
	return p
}

func (p *Printer) SetOutput(writers ...io.Writer) *Printer {
	var w io.Writer
	if l := len(writers); l == 0 {
		return p
	} else if l == 1 {
		w = writers[0]
	} else {
		w = io.MultiWriter(writers...)
	}
	p.mu.Lock()
	p.Output = w
	p.IsTerminal = terminal.IsTerminal(w)
	p.mu.Unlock()
	return p
}

func (p *Printer) EnableDirectOutput() *Printer {
	p.mu.Lock()
	p.DirectOutput = true
	p.mu.Unlock()
	return p
}

func (p *Printer) Print(v interface{}) (int, error) {
	return p.print(v, false)
}

func (p *Printer) Println(v interface{}) (int, error) {
	return p.print(v, true)
}

func (p *Printer) print(v interface{}, appendNewLine bool) (int, error) {
	var (
		b   []byte
		err error
	)
	if p.DirectOutput {
		b, err = p.WriteTo(v, p.Output, appendNewLine)
	} else {
		err = p.Store(v, appendNewLine)
		if err != nil {
			return -1, err
		}
		b, err = p.Flush()
	}
	if len(p.handlers) > 0 {
		res := withValue(v).withErr(err).withContents(b)
		for _, h := range p.handlers {
			h(res)
		}
	}
	return len(b), err
}

func (p *Printer) readAndConsume() ([]byte, error) {
	b, err := ioutil.ReadAll(p.Reader)
	if err != nil && err != io.EOF {
		return b, err
	}
	return b, nil
}

func (p *Printer) Flush() ([]byte, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	b, err := p.readAndConsume()

	if err != nil {
		return nil, err
	}
	_, err = p.Output.Write(b)
	return b, err
}

func (p *Printer) Store(v interface{}, appendNewLine bool) error {
	_, err := p.WriteTo(v, p.Writer, appendNewLine)
	return err
}

func (p *Printer) WriteTo(v interface{}, w io.Writer, appendNewLine bool) ([]byte, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	var marshaler Marshaler
	if m, ok := v.(Marshaled); ok {
		marshaler = fromMarshaled(m)
	} else if m, ok := v.(Marshaler); ok {
		marshaler = m
	} else {
		if p.marshal != nil {
			marshaler = p.marshal
		}
	}
	var (
		b   []byte
		err error
	)
	if hijack := p.hijack; hijack != nil {
		ctx := acquireCtx(v, p)
		defer releaseCtx(ctx)

		hijack(ctx)

		if ctx.canceled {
			return nil, ErrCanceled
		}

		b, err = ctx.marshalResult.b, ctx.marshalResult.err

		if err != nil {
			return b, err
		}
	}

	if len(b) == 0 {
		if marshaler == nil {
			return nil, ErrSkipped
		}

		b, err = marshaler.Marshal(v)
		if err != nil {
			return b, err
		}
	}
	_, err = w.Write(b)
	if appendNewLine && err == nil {
		w.Write(NewLine)
	}
	return b, err
}

func (p *Printer) Hijack(cb func(ctx *Ctx)) *Printer {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.hijack == nil {
		p.hijack = cb
		return p
	}
	oldCb := p.hijack
	newCb := cb
	p.hijack = func(ctx *Ctx) {
		oldCb(ctx)
		if ctx.continueToNext {
			newCb(ctx)
		}
	}
	return p
}

type PrintResult struct {
	Written  int
	Error    error
	Contents []byte
	Value    interface{}
}

func (p PrintResult) IsOK() bool {
	return p.Error == nil && len(p.Contents) > 0
}

func (p PrintResult) IsFailure() bool {
	return !p.IsOK()
}

var printResult = PrintResult{}

func withValue(v interface{}) PrintResult {
	printResult.Value = v
	return printResult
}

func (p PrintResult) withErr(err error) PrintResult {
	if err != nil {
		p.Written = -1
	}
	p.Error = err
	return p
}

func (p PrintResult) withContents(b []byte) PrintResult {
	if p.Error != nil {
		p.Written = -1
	} else {
		p.Written = len(b)
		p.Contents = b
	}
	return p
}

func (p *Printer) Handle(h func(PrintResult)) *Printer {
	p.mu.Lock()
	p.handlers = append(p.handlers, h)
	p.mu.Unlock()
	return p
}

func (p *Printer) restore(b []byte) {
	p.Writer.Write(b)
}

func (p *Printer) Scan(r io.Reader, addNewLine bool) (cancel func()) {
	var canceled uint32
	shouldCancel := func() bool {
		return atomic.LoadUint32(&canceled) > 0
	}
	cancel = func() {
		atomic.StoreUint32(&canceled, 1)
	}

	go func() {
		scanner := bufio.NewScanner(r)

		for {
			if shouldCancel() {
				break
			}
			if scanner.Scan() {
				if shouldCancel() {

					p.restore(scanner.Bytes())
					break
				}
				text := scanner.Bytes()
				if addNewLine {
					text = append(text, NewLine...)
				}
				p.Print(text)
			}
			if err := scanner.Err(); err != nil {

			}
		}
	}()

	return cancel
}
