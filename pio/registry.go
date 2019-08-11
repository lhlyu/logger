package pio

import (
	"errors"
	"io"
	"sort"
	"sync"
)

type Registry struct {
	printers []*Printer
	mu       sync.Mutex
	once     sync.Once
}

func NewRegistry() *Registry {
	return new(Registry)
}

func (reg *Registry) RegisterPrinter(printer *Printer) *Registry {

	if printerName := printer.Name; reg.Get(printerName) != nil {
		reg.Remove(printerName)
	}
	reg.mu.Lock()

	reg.printers = append(reg.printers, printer)
	reg.mu.Unlock()
	return reg
}

func (reg *Registry) Register(printerName string, output io.Writer) *Printer {
	p := NewPrinter(printerName, output)
	reg.RegisterPrinter(p)
	return p
}

func (reg *Registry) Get(printerName string) *Printer {
	reg.mu.Lock()
	defer reg.mu.Unlock()
	for _, p := range reg.printers {
		if p.Name == printerName {
			return p
		}
	}
	return nil
}

func (reg *Registry) Remove(printerName string) *Registry {
	reg.mu.Lock()
	for i, p := range reg.printers {
		if p.Name == printerName {
			reg.printers = append(reg.printers[:i], reg.printers[i+1:]...)
			break
		}
	}
	reg.mu.Unlock()
	return reg
}

func (reg *Registry) Print(v interface{}) (n int, err error) {
	return reg.printAll(v, false)
}

func (reg *Registry) Println(v interface{}) (n int, err error) {
	return reg.printAll(v, true)
}

func (reg *Registry) printAll(v interface{}, appendNewLine bool) (n int, err error) {

	reg.once.Do(func() {
		reg.mu.Lock()
		sort.Slice(reg.printers, func(i, j int) bool {
			return reg.printers[i].priority > reg.printers[j].priority
		})
		reg.mu.Unlock()
	})

	for _, p := range reg.printers {
		prevErr := err

		printFunc := p.Print
		if appendNewLine {
			printFunc = p.Println
		}

		n, err = printFunc(v)

		if !p.Chained && n > 0 {
			break
		}
		n, err = combineOutputResult(n, err, prevErr)
	}
	return
}

func combineOutputResult(n int, err error, prevErr error) (totalN int, totalErr error) {
	if err != nil {
		if prevErr != nil {
			totalErr = errors.New(prevErr.Error() + string(NewLine) + err.Error())
		}
	}

	totalN += n
	return
}

func (reg *Registry) Scan(r io.Reader, addNewLine bool) (cancel func()) {
	lp := len(reg.printers)
	if lp == 0 {
		return func() {}
	}

	cancelFuncs := make([]func(), lp, lp)
	cancel = func() {
		for _, c := range cancelFuncs {
			c()
		}
	}

	for i, p := range reg.printers {
		cancelFuncs[i] = p.Scan(r, addNewLine)
	}

	return cancel
}

func (reg *Registry) restore(b []byte) {
	for _, p := range reg.printers {
		p.restore(b)
	}
}
