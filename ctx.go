package logger

import (
	"context"
	"io"
)

type Ctx struct {
	Out     io.Writer
	Level   Level
	Values  []interface{}
	Context context.Context
	Err     error
	stop    bool
}

func newCtx(logger *Logger, values []interface{}) *Ctx {
	return &Ctx{
		Out:     logger.Out,
		Level:   logger.Level,
		Values:  values,
		Context: logger.Context,
	}
}

func (c *Ctx) Stop() {
	c.stop = true
}
