package tasks

import (
	"encoding/json"

	"github.com/openPanel/core/app/clients/rpc"
	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/manager/router"
	"github.com/openPanel/core/app/tools/convert"
)

func EstimateAndBroadcastLinkState() {
	infos := router.EstimateAndStoreLatencies()

	router.UpdateLinkStates(infos)

	linkStates := convert.LinkStatesRouterToPb(router.GetLinkStates())
	broadcastPayload, err := json.Marshal(linkStates)
	if err != nil {
		log.Errorf("cron: failed to marshal link states: %v", err)
		return
	}

	err = rpc.Broadcast([]rpc.BroadcastMessage{{
		Type:    pb.BroadcastType_NOTIFY_LINK_STATE_CHANGE,
		Payload: string(broadcastPayload),
	}})
	if err != nil {
		log.Errorf("cron: failed to broadcast link state: %v", err)
		return
	}
	log.Infof("cron: broadcast link state periodically")
}
