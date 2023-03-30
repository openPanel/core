package utils

import (
	"os"
	"os/signal"
)

func WaitExit() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan
}
