package rpc

import (
	"context"
	"net/netip"

	"github.com/Rican7/retry"
	"github.com/Rican7/retry/strategy"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/manager/router"
	"github.com/openPanel/core/app/tools/rpcDialer"
	"github.com/openPanel/core/app/tools/utils/netUtils"
)

// TryUpdateRouterNode A node that is not starting for the first time tries to
// load the current cluster information from one of its neighbors
func TryUpdateRouterNode(targets []Target) ([]netip.AddrPort, error) {
	var addrs []netip.AddrPort
	action := func(attempt uint) error {
		currentTarget := targets[attempt]
		conn, err := rpcDialer.DialWithAddress(currentTarget.AddrPort.String(), currentTarget.ServerId)
		if err != nil {
			log.Infof("failed to connect to %s: %s", currentTarget.AddrPort.String(), err)
			return err
		}
		defer func(conn *grpc.ClientConn) {
			_ = conn.Close()
		}(conn)

		client := pb.NewInitializeServiceClient(conn)
		info, err := client.GetClusterInfo(context.Background(), &emptypb.Empty{})
		if err != nil {
			return err
		}

		routerNodes := make([]router.Node, 0, len(info.Nodes)+1)
		routerNodes = append(routerNodes, router.Node{
			Id:       global.App.NodeInfo.ServerId,
			AddrPort: netUtils.NewAddrPortWithIP(global.App.NodeInfo.ServerPublicIP, global.App.NodeInfo.ServerPort),
		})
		for _, node := range info.Nodes {
			routerNodes = append(routerNodes, router.Node{
				Id:       node.Id,
				AddrPort: netUtils.NewAddrPortWithString(node.Ip, int(node.Port)),
			})
		}

		router.SetNodes(routerNodes)

		addrs = make([]netip.AddrPort, len(info.Nodes))
		for i, node := range info.Nodes {
			addrs[i] = netUtils.NewAddrPortWithString(node.Ip, int(node.Port))
		}

		// Test latency

		return nil
	}

	err := retry.Retry(action, strategy.Limit(uint(len(targets))))
	if err != nil {
		return nil, err
	}
	return addrs, nil
}
