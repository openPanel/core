package cron

import (
	"time"

	"github.com/go-co-op/gocron"
)

var scheduler = gocron.NewScheduler(time.UTC)

func Start() {
	scheduler.
		EveryRandom(10, 20).
		Minute().
		StartImmediately()

	scheduler.Len()
}
