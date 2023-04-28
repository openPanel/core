package rpc

import (
	"context"
	"net/netip"
	"sync"

	"github.com/Rican7/retry"
	"github.com/Rican7/retry/strategy"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/openPanel/core/app/clients/tcp"
	"github.com/openPanel/core/app/db/repo/shared"
	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/manager/router"
	"github.com/openPanel/core/app/tools/rpcDialer"
	"github.com/openPanel/core/app/tools/utils/netUtils"
)

func syncLinkStates(target Target) ([]netip.AddrPort, error) {
	conn, err := rpcDialer.DialWithAddress(target.AddrPort.String(), target.ServerId)
	if err != nil {
		log.Infof("failed to connect to %s: %s", target.AddrPort.String(), err)
		return nil, err
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

	addrs := make([]netip.AddrPort, len(info.Nodes))
	for i, node := range info.Nodes {
		addrs[i] = netUtils.NewAddrPortWithString(node.Ip, int(node.Port))
	}

	// Test latency
	// TODO: implement me

	return nil
}

// SyncLinkStates A node that is not starting the first time tries to
// load the current cluster information from one of its neighbors
// Called before the node connected to the cluster, thus should not use dqlite repo
func SyncLinkStates(targets []Target) ([]netip.AddrPort, error) {
	var addrs []netip.AddrPort
	action := func(attempt uint) error {
		currentTarget := targets[attempt]

		var err error
		addrs, err = syncLinkStates(currentTarget)
		return err
	}

	err := retry.Retry(action, strategy.Limit(uint(len(targets))))
	if err != nil {
		return nil, err
	}
	return addrs, nil
}

// CollectLatencies works in the cluster, test latency to new node, everything configured
func CollectLatencies(ctx context.Context, ip string, port int, id string) (router.LinkStates, error) {
	nodes, err := shared.NodeRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	wg := sync.WaitGroup{}
	wg.Add(len(nodes))
	lst := router.LinkStates{}
	lock := sync.Mutex{}

	updateLst := func(curId string, latency int) {
		lock.Lock()
		defer lock.Unlock()

		lst[router.Edge{
			From: curId,
			To:   id,
		}] = latency
	}

	for _, node := range nodes {
		go func(curId string, addr netip.AddrPort) {
			defer wg.Done()

			if curId == global.App.NodeInfo.ServerId {
				latency, err := tcp.Ping(netUtils.NewAddrPortWithString(ip, port))
				if err != nil {
					log.Infof("failed to tcp ping %s: %s", addr.String(), err)
					return
				}
				updateLst(curId, latency)
			} else {
				conn, err := rpcDialer.DialWithServerId(curId)
				if err != nil {
					log.Errorf("failed to connect to %s: %v", curId, err)
					return
				}

				client := pb.NewInitializeServiceClient(conn)
				resp, err := client.EstimateLatency(ctx, &pb.EstimateLatencyRequest{
					Ip:   addr.Addr().String(),
					Port: int32(addr.Port()),
				})
				if err != nil {
					log.Infof("failed to estimate latency to %s: %s", addr.String(), err)
					return
				}

				updateLst(curId, int(resp.Latency))
			}
		}(node.ID, netUtils.NewAddrPortWithString(node.IP, node.Port))
	}

	wg.Wait()

	return lst, nil
}
