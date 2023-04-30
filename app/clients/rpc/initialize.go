package rpc

import (
	"context"

	"github.com/Rican7/retry"
	"github.com/Rican7/retry/strategy"
	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/manager/router"
	"github.com/openPanel/core/app/tools/convert"
	"github.com/openPanel/core/app/tools/rpcDialer"
)

func tryConnectToNeighbor(targets []Target) (pb.InitializeServiceClient, error) {
	var client pb.InitializeServiceClient
	action := func(attempt uint) error {
		target := targets[attempt]

		conn, err := rpcDialer.DialWithAddress(target.AddrPort.String(), target.ServerId)
		if err != nil {
			log.Infof("failed to connect to %s: %s", target.AddrPort.String(), err)
			return err
		}

		client = pb.NewInitializeServiceClient(conn)
		return nil
	}

	err := retry.Retry(action, strategy.Limit(uint(len(targets))))
	if err != nil {
		return nil, err
	}

	return client, nil
}

// SyncNodesAndLinkStates A node that is not starting the first time tries to
// load the current cluster information from one of its neighbors
// Called before the node connected to the cluster, thus should not use dqlite repo
func SyncNodesAndLinkStates(targets []Target) ([]router.Node, router.LinkStates, error) {
	client, err := tryConnectToNeighbor(targets)
	if err != nil {
		return nil, nil, err
	}

	// during resume, nodes already contains the node itself
	info, err := client.GetClusterInfo(context.Background(), &emptypb.Empty{})
	if err != nil {
		return nil, nil, err
	}

	routerNodes := convert.NodesPbToRouter(info.Nodes)

	otherNodes := lo.Filter(routerNodes, func(node router.Node, _ int) bool {
		return node.Id != global.App.NodeInfo.ServerId
	})

	lstToOtherNodes := router.EstimateLatencies(otherNodes, global.App.NodeInfo.ServerId)

	// sync link states
	fullLinkStates, err := client.UpdateLinkState(context.Background(), &pb.UpdateLinkStateRequest{
		LinkStates: convert.LinkStatesRouterToPb(lstToOtherNodes),
	})
	if err != nil {
		return nil, nil, err
	}

	return routerNodes, convert.LinkStatesPbToRouter(fullLinkStates.LinkStates), nil
}
