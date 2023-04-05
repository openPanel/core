package cron

import (
	"time"

	"github.com/go-co-op/gocron"

	"github.com/openPanel/core/app/bootstrap/clean"
	"github.com/openPanel/core/app/global/log"
)

var scheduler = gocron.NewScheduler(time.UTC)

func Start() {
	registerDefaultCronTasks()
	go scheduler.StartBlocking()

	clean.RegisterCleanup(func() {
		log.Info("cron manager: stopping scheduler")
		scheduler.Stop()
		log.Info("cron manager: scheduler stopped")

	})
}

func Op(fn func(s *gocron.Scheduler)) {
	fn(scheduler)
}

var DefaultCronTasks []func(s *gocron.Scheduler)

func registerDefaultCronTasks() {
	for _, task := range DefaultCronTasks {
		task(scheduler)
	}
}
