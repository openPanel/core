// Package stop detect exit signal and execute cleanup function
package stop

import (
	"os"
	"os/signal"
	"sync"
	"time"

	"golang.org/x/sys/unix"

	"github.com/openPanel/core/app/constant"
	"github.com/openPanel/core/app/global/log"
)

var cleanLock sync.Mutex

var cleanOps map[constant.StopID]cleanOp

type cleanOp struct {
	// deps is the dependencies of this cleanup operation
	// if any of the dependencies is not executed, this cleanup operation will not be executed
	deps []constant.StopID
	// fn is the cleanup function
	fn func()
}

func RegisterCleanup(cleanup func(), id constant.StopID, deps ...constant.StopID) {
	ret := cleanLock.TryLock()
	if !ret {
		log.Panicf("RegisterCleanup should be called in linear order, but it's not")
	}

	defer cleanLock.Unlock()
	cleanOps[id] = cleanOp{
		deps: deps,
		fn:   cleanup,
	}
}

func buildExecuteOrder() {
	inDeg := map[constant.StopID]int{}

	for id := range cleanOps {
		inDeg[id] = 0
	}

	for id := range cleanOps {
		for _, dep := range cleanOps[id].deps {
			inDeg[dep]++
		}
	}

	var q []constant.StopID
	for id := range cleanOps {
		if inDeg[id] == 0 {
			q = append(q, id)
		}
	}

}

func RunEndless() {
	cleanLock.Lock()
	defer cleanLock.Unlock()

	ch := make(chan os.Signal, 3)

	signal.Notify(ch, unix.SIGPWR, unix.SIGINT, unix.SIGQUIT, unix.SIGTERM)

	sig := <-ch

	go func() {
		for sig := range ch {
			log.Warnf("Received signal %s while cleaning up, ignore", sig.String())
		}
	}()

	go func() {
		time.Sleep(8 * time.Second)
		log.Panicf("Timed out while cleaning up, exiting")
	}()

	log.Infof("Received signal %s, cleaning up", sig.String())

	os.Exit(0)
}
