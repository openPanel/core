package services

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/openPanel/core/app/db/repo/shared"
	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/manager/router"
	"github.com/openPanel/core/app/tools/utils/netUtils"
)

var LinkStateService pb.LinkStateServiceServer = new(linkStateService)

type linkStateService struct{}

func (l linkStateService) UpdateLinkState(_ context.Context, request *pb.LinkStateUpdateRequest) (*emptypb.Empty, error) {
	infos := map[router.Edge]int{}
	for _, state := range request.LinkState {
		edge := router.Edge{
			From: state.From,
			To:   state.To,
		}
		infos[edge] = int(state.Latency)
	}
	router.UpdateRouterInfo(infos)
	return &emptypb.Empty{}, nil
}

func (l linkStateService) NotifyNodeUpdate(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	eNodes, err := shared.NodeRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	routerNodes := make([]router.Node, len(eNodes))
	for i, node := range eNodes {
		routerNodes[i] = router.Node{
			Id:       node.ID,
			AddrPort: netUtils.NewAddPortWithString(node.IP, node.Port),
		}
		router.SetNodes(routerNodes)
	}
	return &emptypb.Empty{}, nil
}
