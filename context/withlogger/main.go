// Command withlogger exmaple how to use context and logger using `WithContext` and `FromContext` function.
package main

import (
	"context"

	"go.uber.org/zap"

	"github.com/zchee/go-sandbox/context/withlogger/logger"
	"github.com/zchee/go-sandbox/context/withlogger/logger/wrapper"
)

var exampleField = zap.String("key", "value")

func main() {
	// NOTE(zchee): You can use context cancelling with following snippet
	//
	//  ctx, cancel := context.WithCancel(context.Background())
	//  defer cancel() // cancelling context after the end of main function.

	log := logger.NewLogger()
	log.Info("first log", exampleField)

	jsonLogger := logger.NewJsonLogger()
	jsonLogger.Info("second log with json payload style", exampleField)

	// global context.
	ctx := context.Background()

	// if comment out, will panic because not have logger.Logger value into ctx.
	// logger.FromContext(ctx).Info("third log with logger.FromContext", exampleField)

	// embedded logger.Logger value into ctx.
	ctx = logger.WithContext(ctx)

	// use embedded logger.Logger context value with logger.FromContext function
	logger.FromContext(ctx).Info("third log with logger.FromContext", exampleField)

	OutputFoo(ctx)
	OutputWrappedFoo(ctx)
}

// OutputFoo outputs foo log.
func OutputFoo(ctx context.Context) {
	logger.FromContext(ctx).Info("Fooooooooo!", exampleField)
}

// OutputWrappedFooLog outputs wrapped foo log.
//
// ctx arg must be have logger.Logger value.
func OutputWrappedFoo(ctx context.Context) {
	wrapper.NewWrapperLogger(ctx).Info("wrapped Fooooooooo!", exampleField)
}
