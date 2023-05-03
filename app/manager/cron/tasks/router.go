package tasks

import (
	"github.com/openPanel/core/app/clients/rpc"
	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/manager/router"
	"github.com/openPanel/core/app/tools/broadcast"
)

func EstimateAndBroadcastLinkState() {
	infos := router.EstimateAndStoreLatencies()

	router.UpdateLinkStates(infos)

	payload, err := broadcast.GetRouterPayload(infos, nil, nil, nil)

	err = rpc.Broadcast(rpc.BroadcastMessage{
		Type:    pb.BroadcastType_LINK_STATE_CHANGE,
		Payload: payload,
	})
	if err != nil {
		log.Errorf("cron: failed to broadcast link state: %v", err)
		return
	}
	log.Infof("cron: broadcast link state periodic")
}
