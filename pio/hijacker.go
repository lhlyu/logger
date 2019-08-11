package pio

import (
	"errors"
	"sync"
)

type Hijacker func(*Ctx)

var (
	ErrCanceled = errors.New("canceled")

	ErrSkipped = errors.New("skipped")
)

var cPool = sync.Pool{New: func() interface{} { return &Ctx{} }}

func acquireCtx(v interface{}, printer *Printer) *Ctx {
	ctx := cPool.Get().(*Ctx)
	ctx.Printer = printer
	ctx.Value = v

	ctx.marshalResult.b = ctx.marshalResult.b[0:0]
	ctx.marshalResult.err = nil
	ctx.canceled = false
	ctx.continueToNext = false
	return ctx
}

func releaseCtx(ctx *Ctx) {
	cPool.Put(ctx)
}

type Ctx struct {
	Printer       *Printer
	Value         interface{}
	marshalResult struct {
		b   []byte
		err error
	}
	continueToNext bool
	canceled       bool
}

func (ctx *Ctx) MarshalValue() ([]byte, error) {
	if len(ctx.marshalResult.b) > 0 {
		return ctx.marshalResult.b, ctx.marshalResult.err
	}

	if ctx.Printer.marshal == nil {
		return nil, ErrSkipped
	}

	b, err := ctx.Printer.marshal(ctx.Value)
	ctx.marshalResult.b = b
	ctx.marshalResult.err = err
	return b, err
}

func (ctx *Ctx) Store(result []byte, err error) {
	ctx.marshalResult.b = result
	ctx.marshalResult.err = err
}

func (ctx *Ctx) Cancel() {
	ctx.canceled = true
}

func (ctx *Ctx) Next() {
	ctx.continueToNext = true
}
