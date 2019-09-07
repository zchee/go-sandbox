// Package logger provide a how to use context and logger(in this case, zap.Logger) example.
package logger

import (
	"context"

	"go.uber.org/zap"
)

// Logger represents a zap.Logger.
type Logger struct {
	*zap.Logger
}

// NewLogger returns the new Logger for console.
func NewLogger() *Logger {
	cfg := zap.NewDevelopmentConfig()

	l, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return &Logger{
		Logger: l,
	}
}

// NewJsonLogger returns the new Logger with outputs json payload style.
func NewJsonLogger() *Logger {
	cfg := zap.NewDevelopmentConfig()
	cfg.Encoding = "json"

	l, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return &Logger{
		Logger: l,
	}
}

type ctxKey struct{} // empty, `struct{}` no allocate memory

// WithContext returns context.Context with embedded Logger.
//
// In `ctxKey{}` second arg, *DO NOT* use the pointer.
// Pass by *value* (means 値渡し) for the avoid unnecessary memory allocation.
func WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxKey{}, NewLogger() /* In production, pass to actual your Logger struct */)
}

// FromContext extracts and returns the Logger from ctx value.
//
// No needs:
//  log, ok := ctx.Value(ctxKey{}).(*Logger)
//  if !ok { panic(...) }
func FromContext(ctx context.Context) *Logger {
	return ctx.Value(ctxKey{}).(*Logger)
}

// ----------------------------------------------------------------------------
// Advance Usage: how to use pointer context Value key.
//
// See logger/wrapper package.

// Key represents a global logger package context key.
//
// DO NOT USE IT your package implementation. For internal use.
var Key = &ctxKey{}
