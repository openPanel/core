package cron

import (
	"time"

	"github.com/go-co-op/gocron"

	"github.com/openPanel/core/app/constant"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/manager/detector/stop"
)

var scheduler = gocron.NewScheduler(time.UTC)

func Start() {
	registerDefaultCronTasks()

	stop.RegisterCleanup(func() {
		scheduler.Stop()
		log.Info("cron manager: scheduler stopped")
	}, constant.StopIDCron, constant.StopIDLogger)

	scheduler.StartAsync()
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
