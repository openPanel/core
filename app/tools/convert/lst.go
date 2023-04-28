package convert

import (
	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/manager/router"
)

func LinkStatesRouterToPb(lst router.LinkStates) []*pb.LinkState {
	var pbLst []*pb.LinkState
	for edge, latency := range lst {
		pbLst = append(pbLst, &pb.LinkState{
			From:    edge.From,
			To:      edge.To,
			Latency: int32(latency),
		})
	}
	return pbLst
}

func LinkStatesPbToRouter(lst []*pb.LinkState) router.LinkStates {
	var routerLst = make(router.LinkStates)
	for _, linkState := range lst {
		routerLst[router.Edge{
			From: linkState.From,
			To:   linkState.To,
		}] = int(linkState.Latency)
	}
	return routerLst
}
