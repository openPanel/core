package convert

import (
	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/manager/router"
)

func LinkStatesRouterToPb(lst router.LinkStates) []*pb.LinkState {
	var pbLst = make([]*pb.LinkState, len(lst))
	for i, linkState := range lst {
		pbLst[i] = &pb.LinkState{
			From:    linkState.From,
			To:      linkState.To,
			Latency: int32(linkState.Latency),
		}
	}
	return pbLst
}

func LinkStatesPbToRouter(lst []*pb.LinkState) router.LinkStates {
	var routerLst = make(router.LinkStates, len(lst))
	for i, linkState := range lst {
		routerLst[i] = router.LinkState{
			From:    linkState.From,
			To:      linkState.To,
			Latency: int(linkState.Latency),
		}
	}
	return routerLst
}

func LinkStatesMerge(lst ...router.LinkStates) router.LinkStates {
	size := 0
	for _, l := range lst {
		size += len(l)
	}

	var mergedLst = make(router.LinkStates, 0, size)
	for _, l := range lst {
		mergedLst = append(mergedLst, l...)
	}
	return mergedLst
}
