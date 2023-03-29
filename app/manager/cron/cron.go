package cron

import (
	"time"

	"github.com/go-co-op/gocron"
)

var scheduler = gocron.NewScheduler(time.UTC)

func Start() {
	go scheduler.StartBlocking()
}

func Op(fn func(s *gocron.Scheduler)) {
	fn(scheduler)
}
