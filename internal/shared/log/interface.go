package log

import "context"

var Logger *ZapLogger

type Log interface {
	FromCtx(ctx context.Context) *ZapLogger
	WithCtx(ctx context.Context) context.Context
	Debug(msg string, fields ...interface{})
	Info(msg string, fields ...interface{})
	Warn(msg string, fields ...interface{})
	Error(msg string, fields ...interface{})
	Panic(msg string, fields ...interface{})
	DPanic(msg string, fields ...interface{})
	Fatal(msg string, fields ...interface{})
}
