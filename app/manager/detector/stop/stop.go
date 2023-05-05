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

var cleanups []func()
var cleanLock sync.Mutex

func RegisterCleanup(cleanup func(), id constant.StopID, deps ...constant.StopID) {
	ret := cleanLock.TryLock()
	if !ret {
		log.Panicf("RegisterCleanup should be called in linear order, but it's not")
	}

	defer cleanLock.Unlock()
	cleanups = append(cleanups, cleanup)
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

	wg := sync.WaitGroup{}
	for _, cleanup := range cleanups {
		wg.Add(1)
		go func(cleanup func()) {
			defer wg.Done()
			cleanup()
		}(cleanup)
	}
	wg.Wait()

	log.Info("Cleaned up, exiting")
	_ = log.Sync()

	os.Exit(0)
}
