package log

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var once sync.Once

type ctxKey struct{}

type ZapLogger struct {
	logger *zap.Logger
}

// Get initializes a zap.Logger instance if it has not been initialized
// already and returns the same instance for subsequent calls.
func Init() {
	var logger *zap.Logger

	stdout := zapcore.AddSync(os.Stdout)

	level := zap.InfoLevel
	levelEnv := os.Getenv("LOG_LEVEL")

	// If the LOG_LEVEL environment variable is set, parse it and set the log level
	if levelEnv != "" {
		levelFromEnv, err := zapcore.ParseLevel(levelEnv)
		if err != nil {
			log.Println(
				fmt.Errorf("invalid level, defaulting to INFO: %w", err),
			)
		}

		level = levelFromEnv
	}

	logLevel := zap.NewAtomicLevelAt(level)

	var consoleEncoder zapcore.Encoder

	// If the APP_ENV environment variable is set to "prod", use the production config
	// else use the development config
	if os.Getenv("APP_ENV") == "prod" {
		productionCfg := zap.NewProductionEncoderConfig()
		productionCfg.TimeKey = "timestamp"
		productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder

		consoleEncoder = zapcore.NewConsoleEncoder(productionCfg)
	} else {
		developmentCfg := zap.NewDevelopmentEncoderConfig()
		developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

		consoleEncoder = zapcore.NewConsoleEncoder(developmentCfg)
	}

	var gitRevision string

	buildInfo, ok := debug.ReadBuildInfo()
	if ok {
		for _, v := range buildInfo.Settings {
			if v.Key == "vcs.revision" {
				gitRevision = v.Value
				break
			}
		}
	}

	// log to stdout
	core := zapcore.NewCore(consoleEncoder, stdout, logLevel).
		With(
			[]zapcore.Field{
				zap.String("git_revision", gitRevision),
				zap.String("go_version", buildInfo.GoVersion),
			},
		)

	logger = zap.New(core)
	Logger = &ZapLogger{logger: logger}
}

// FromCtx returns the Logger associated with the ctx. If no logger
// is associated, the default logger is returned, unless it is nil
// in which case a disabled logger is returned.
func (l *ZapLogger) FromCtx(ctx context.Context) *ZapLogger {
	// Check if the logger is already attached to the context
	// If it is, return the logger
	if l, ok := ctx.Value(ctxKey{}).(*ZapLogger); ok {
		return l
	}

	// If the logger is not attached to the context, return the no-op logger
	return &ZapLogger{logger: zap.NewNop()}
}

// WithCtx returns a copy of ctx with the Logger attached.
func (l *ZapLogger) WithCtx(ctx context.Context) context.Context {
	if lp, ok := ctx.Value(ctxKey{}).(*ZapLogger); ok {
		// If the logger is already attached to the context, return the context as it is
		if lp == l {
			// Do not store same logger.
			return ctx
		}
	}

	// if the context does not have a logger attached,
	// attach the logger to the context
	return context.WithValue(ctx, ctxKey{}, l)
}

// Debug logs an error message with the given fields.
func (l *ZapLogger) Debug(msg string, fields ...interface{}) {
	l.logger.Sugar().Debugw(msg, fields...)
}

// Info logs an error message with the given fields.
func (l *ZapLogger) Info(msg string, fields ...interface{}) {
	l.logger.Sugar().Infow(msg, fields...)
}

// Warn logs an error message with the given fields.
func (l *ZapLogger) Warn(msg string, fields ...interface{}) {
	l.logger.Sugar().Warnw(msg, fields...)
}

// Error logs an error message with the given fields.
func (l *ZapLogger) Error(msg string, fields ...interface{}) {
	l.logger.Sugar().Errorw(msg, fields...)
}

// Panic logs an error message with the given fields.
func (l *ZapLogger) Panic(msg string, fields ...interface{}) {
	l.logger.Sugar().Panicw(msg, fields...)
}

// DPanic logs an error message with the given fields.
func (l *ZapLogger) DPanic(msg string, fields ...interface{}) {
	l.logger.Sugar().DPanicw(msg, fields...)
}

// Fatal logs an error message with the given fields.
func (l *ZapLogger) Fatal(msg string, fields ...interface{}) {
	l.logger.Sugar().Fatalw(msg, fields...)
}
