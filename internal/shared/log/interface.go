package log

import "context"

var BaseLogger Log

type Log interface {
	FromCtx(ctx context.Context) Log
	WithCtx(ctx context.Context) context.Context
	WithFields(fields ...any) Log
	Debug(msg string, fields ...any)
	Info(msg string, fields ...any)
	Warn(msg string, fields ...any)
	Error(msg string, fields ...any)
	Panic(msg string, fields ...any)
	DPanic(msg string, fields ...any)
	Fatal(msg string, fields ...any)
}
