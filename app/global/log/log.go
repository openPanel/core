package log

import (
	"go.uber.org/zap"

	"github.com/openPanel/core/app/global"
)

var (
	Named func(name string) *zap.SugaredLogger
	With  func(args ...any) *zap.SugaredLogger

	Debug  func(args ...any)
	Info   func(args ...any)
	Warn   func(args ...any)
	Error  func(args ...any)
	DPanic func(args ...any)
	Panic  func(args ...any)
	Fatal  func(args ...any)

	Debugf  func(template string, args ...any)
	Infof   func(template string, args ...any)
	Warnf   func(template string, args ...any)
	Errorf  func(template string, args ...any)
	DPanicf func(template string, args ...any)
	Panicf  func(template string, args ...any)
	Fatalf  func(template string, args ...any)

	Debugw  func(msg string, keysAndValues ...any)
	Infow   func(msg string, keysAndValues ...any)
	Warnw   func(msg string, keysAndValues ...any)
	Errorw  func(msg string, keysAndValues ...any)
	DPanicw func(msg string, keysAndValues ...any)
	Panicw  func(msg string, keysAndValues ...any)
	Fatalw  func(msg string, keysAndValues ...any)

	Debugln  func(args ...any)
	Infoln   func(args ...any)
	Warnln   func(args ...any)
	Errorln  func(args ...any)
	DPanicln func(args ...any)
	Panicln  func(args ...any)
	Fatalln  func(args ...any)

	Sync func() error
)

// UpdateLogger helpful for late binding
func UpdateLogger(logger *zap.SugaredLogger) {
	global.App.Logger = logger.Desugar()

	Named = logger.Named
	With = logger.With

	Debug = logger.Debug
	Info = logger.Info
	Warn = logger.Warn
	Error = logger.Error
	DPanic = logger.DPanic
	Panic = logger.Panic
	Fatal = logger.Fatal

	Debugf = logger.Debugf
	Infof = logger.Infof
	Warnf = logger.Warnf
	Errorf = logger.Errorf
	DPanicf = logger.DPanicf
	Panicf = logger.Panicf
	Fatalf = logger.Fatalf

	Debugw = logger.Debugw
	Infow = logger.Infow
	Warnw = logger.Warnw
	Errorw = logger.Errorw
	DPanicw = logger.DPanicw
	Panicw = logger.Panicw
	Fatalw = logger.Fatalw

	Debugln = logger.Debugln
	Infoln = logger.Infoln
	Warnln = logger.Warnln
	Errorln = logger.Errorln
	DPanicln = logger.DPanicln
	Panicln = logger.Panicln
	Fatalln = logger.Fatalln

	Sync = logger.Sync
}
