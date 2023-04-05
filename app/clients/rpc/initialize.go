package rpc

import (
	"context"
	"net/netip"

	"github.com/Rican7/retry"
	"github.com/Rican7/retry/strategy"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/openPanel/core/app/config"
	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/manager/router"
	"github.com/openPanel/core/app/tools/rpc"
)

// TryUpdateRouterInfo A node that is not starting for the first time tries to
// load the current cluster information from one of its neighbors
func TryUpdateRouterInfo(targets []Target) ([]netip.AddrPort, error) {
	var addrs []netip.AddrPort
	action := func(attempt uint) error {
		currentTarget := targets[attempt]
		conn, err := rpc.DialWithAddress(currentTarget.AddrPort.String(), currentTarget.ServerId)
		if err != nil {
			log.Infof("failed to connect to %s: %s", currentTarget.AddrPort.String(), err)
			return err
		}
		defer func(conn *grpc.ClientConn) {
			_ = conn.Close()
		}(conn)

		client := pb.NewInitializeServiceClient(conn)
		info, err := client.GetNodesInfo(context.Background(), &emptypb.Empty{})
		if err != nil {
			return err
		}

		var cache = make([]router.Node, 0, len(info.Nodes))
		for _, node := range info.Nodes {
			cache = append(cache, router.Node{
				Id:       node.Id,
				AddrPort: netip.AddrPortFrom(netip.MustParseAddr(node.Ip), uint16(node.Port)),
			})
		}

		err = config.UpdateNodesCache(cache)
		if err != nil {
			return err
		}

		routerInfo := map[router.Edge]int{}
		for _, ls := range info.LinkStates {
			routerInfo[router.Edge{
				From: ls.From,
				To:   ls.To,
			}] = int(ls.Latency)
		}
		router.UpdateRouterInfo(routerInfo)

		addrs = make([]netip.AddrPort, len(info.Nodes))
		for i, node := range info.Nodes {
			addrs[i] = netip.AddrPortFrom(netip.MustParseAddr(node.Ip), uint16(node.Port))
		}

		return nil
	}

	err := retry.Retry(action, strategy.Limit(uint(len(targets))))
	if err != nil {
		return nil, err
	}
	return addrs, nil
}
