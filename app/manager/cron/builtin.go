package cron

import (
	"log"

	"github.com/go-co-op/gocron"

	"github.com/openPanel/core/app/manager/cron/tasks"
)

func init() {
	DefaultCronTasks = []Task{
		func(s *gocron.Scheduler) {
			// random interval between 15 and 30 minutes to prevent all nodes from doing the same thing at the same time
			_, err := s.EveryRandom(15, 30).Minutes().WaitForSchedule().Do(tasks.EstimateAndBroadcastLinkState)
			if err != nil {
				log.Fatalf("cron: failed to register task: %v", err)
			}
		},
	}
}
