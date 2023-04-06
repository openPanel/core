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

	clean.RegisterCleanup(func() {
		scheduler.Stop()
		log.Info("cron manager: scheduler stopped")

	})

	go scheduler.StartBlocking()
}

func Op(fn func(s *gocron.Scheduler)) {
	fn(scheduler)
}

type Task func(s *gocron.Scheduler)

var DefaultCronTasks []Task

func registerDefaultCronTasks() {
	for _, task := range DefaultCronTasks {
		task(scheduler)
	}
}
