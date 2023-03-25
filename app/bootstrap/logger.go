package bootstrap

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/openPanel/core/app/global"
	customLog "github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/tools/fileUtils"
)

func initLogger() {
	var c zapcore.EncoderConfig

	if global.IsDev() {
		c = zap.NewDevelopmentEncoderConfig()
	} else {
		c = zap.NewProductionEncoderConfig()
	}

	logFile, err := fileUtils.RequireDataFile("core.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fileEncoder := zapcore.NewJSONEncoder(c)

	c.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(c)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(logFile), zapcore.InfoLevel),
	)

	logger := zap.New(core, zap.AddStacktrace(zapcore.WarnLevel), zap.AddCallerSkip(1))
	if global.IsDev() {
		logger = logger.WithOptions(zap.Development())
	}

	customLog.UpdateLogger(logger.Sugar())
}
