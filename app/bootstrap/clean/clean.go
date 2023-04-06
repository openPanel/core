package clean

import (
	"os"
	"os/signal"
	"sync"
	"time"

	"golang.org/x/sys/unix"

	"github.com/openPanel/core/app/global/log"
)

var cleanups []func()
var cleanLock sync.Mutex

func RegisterCleanup(cleanup func()) {
	cleanLock.Lock()
	defer cleanLock.Unlock()
	cleanups = append(cleanups, cleanup)
}

func WaitClean() {
	ch := make(chan os.Signal, 32)

	signal.Notify(ch, unix.SIGPWR, unix.SIGINT, unix.SIGQUIT, unix.SIGTERM)

	sig := <-ch

	go func() {
		for sig := range ch {
			log.Warnf("Received signal %s while cleaning up, ignore", sig.String())
		}
	}()

	go func() {
		time.Sleep(15 * time.Second)
		log.Fatalf("Timed out while cleaning up, exiting")
	}()

	log.Infof("Received signal %s, cleaning up", sig.String())

	cleanLock.Lock()
	defer cleanLock.Unlock()

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
