// Package wrapper wrapped logger package with zap.String field.
package wrapper

import (
	"context"

	"go.uber.org/zap"

	"github.com/zchee/go-sandbox/context/withlogger/logger"
)

const keyname = "wrapper"

// WrapLogger represents a wrap logger.
type WrapLogger struct {
	*zap.Logger
}

// NewWrapperLogger returns the new WrapLogger with wrapped logger.Logger.
//
// Example for advance usage of logger.Key export context key.
func NewWrapperLogger(ctx context.Context) *WrapLogger {
	log := ctx.Value(*logger.Key).(*logger.Logger)
	wrapped := log.With(zap.String("key", keyname))

	return &WrapLogger{
		Logger: wrapped,
	}
}
