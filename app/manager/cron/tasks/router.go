package tasks

import (
	"context"
	"sync"

	"github.com/openPanel/core/app/clients/rpc"
	"github.com/openPanel/core/app/db/repo/shared"
	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/manager/router"
)

func EstimateAndBroadcastLinkState() {
	infos := router.EstimateAndStoreLatencies()
	nodes, err := shared.NodeRepo.GetAll(context.Background())
	if err != nil {
		log.Errorf("cron: failed to load nodes cache: %v", err)
	}

	linkStates := make([]*pb.LinkState, 0, len(infos))
	for edge, latency := range infos {
		linkStates = append(linkStates, &pb.LinkState{
			From:    edge.From,
			To:      edge.To,
			Latency: int32(latency),
		})
	}

	errs := make([]error, 0, len(nodes))

	wg := sync.WaitGroup{}
	wg.Add(len(nodes))
	for _, node := range nodes {
		go func(target string) {
			defer wg.Done()
			err := rpc.UpdateLinkState(target, linkStates)
			if err != nil {
				log.Warnf("cron: failed to broadcast link state to node %s: %v", target, err)
				errs = append(errs, err)
			}
		}(node.ID)
	}

	wg.Wait()
	if len(errs) > len(nodes)/2 {
		log.Errorf("cron: failed to broadcast link state to more than half of nodes")
	}
	log.Infof("cron: broadcast link state to %d nodes", len(nodes)-len(errs))
}
