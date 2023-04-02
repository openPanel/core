package dqlite

import (
	"github.com/canonical/go-dqlite/client"

	"github.com/openPanel/core/app/global/log"
)

func getDqliteLogger() client.LogFunc {
	namedLogger := log.Named("dqlite")
	return func(l client.LogLevel, format string, a ...any) {
		switch l {
		case client.LogNone:
			return
		case client.LogDebug:
			namedLogger.Debugf(format, a...)
		case client.LogInfo:
			namedLogger.Infof(format, a...)
		case client.LogWarn:
			namedLogger.Warnf(format, a...)
		case client.LogError:
			namedLogger.Errorf(format, a...)
		}
	}
}
