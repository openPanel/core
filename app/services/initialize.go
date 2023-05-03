package services

import (
	"context"
	"net/netip"
	"sync"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/openPanel/core/app/clients/rpc"
	"github.com/openPanel/core/app/clients/tcp"
	. "github.com/openPanel/core/app/db/repo/shared"
	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/manager/router"
	"github.com/openPanel/core/app/tools/broadcast"
	"github.com/openPanel/core/app/tools/ca"
	"github.com/openPanel/core/app/tools/convert"
	"github.com/openPanel/core/app/tools/rpcDialer"
	"github.com/openPanel/core/app/tools/utils/netUtils"
)

var InitializeService pb.InitializeServiceServer = new(initializeService)

type initializeService struct{}

func (s *initializeService) UpdateLinkState(ctx context.Context, request *pb.UpdateLinkStateRequest) (*pb.UpdateLinkStateResponse, error) {
	lstFromNewNode := convert.LinkStatesPbToRouter(request.LinkStates)
	lstToNewNode, err := collectLatencies(ctx, request.Ip, int(request.Port), request.ServerID)
	if err != nil {
		return nil, err
	}

	router.UpdateLinkStates(lstFromNewNode, lstToNewNode)

	payload, err := broadcast.GetRouterPayload(convert.LinkStatesMerge(lstFromNewNode, lstToNewNode), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	// broadcast new link state
	err = rpc.Broadcast(rpc.BroadcastMessage{
		Type:    pb.BroadcastType_LINK_STATE_CHANGE,
		Payload: payload,
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateLinkStateResponse{
		LinkStates: convert.LinkStatesRouterToPb(router.GetLinkStates()),
	}, nil
}

func (s *initializeService) EstimateLatency(_ context.Context, request *pb.EstimateLatencyRequest) (*pb.EstimateLatencyResponse, error) {
	latency, err := tcp.Ping(netUtils.NewAddrPortWithString(request.Ip, int(request.Port)))
	if err != nil {
		return nil, err
	}

	return &pb.EstimateLatencyResponse{
		Latency: int32(latency),
	}, nil
}

func (s *initializeService) GetClusterInfo(ctx context.Context, _ *emptypb.Empty) (*pb.GetClusterInfoResponse, error) {
	resp := new(pb.GetClusterInfoResponse)

	nodes, err := NodeRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	resp.Nodes = convert.NodesDbToPb(nodes)

	return resp, nil
}

func (s *initializeService) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	clientCert, err := ca.SignCsr(global.App.NodeInfo.ClusterCaCert, global.App.ClusterInfo.CaKey, request.Csr)
	if err != nil {
		return nil, err
	}

	lstFromNewNode := convert.LinkStatesPbToRouter(request.LinkStates)
	lstToNewNode, err := collectLatencies(ctx, request.Ip, int(request.Port), request.ServerID)
	if err != nil {
		return nil, err
	}

	newRouterNode := router.Node{
		Id:       request.ServerID,
		AddrPort: netUtils.NewAddrPortWithString(request.Ip, int(request.Port)),
	}
	router.AddNodes([]router.Node{newRouterNode})
	router.UpdateLinkStates(lstFromNewNode, lstToNewNode)

	payload, err := broadcast.GetRouterPayload(convert.LinkStatesMerge(lstFromNewNode, lstToNewNode), &[]router.Node{
		newRouterNode,
	}, nil, nil)
	if err != nil {
		return nil, err
	}

	err = rpc.Broadcast(rpc.BroadcastMessage{
		Type:    pb.BroadcastType_LINK_STATE_CHANGE,
		Payload: payload,
	})
	if err != nil {
		return nil, err
	}

	// we have to add node after we update link states, because we can not connect to new node at this time
	err = NodeRepo.AddNode(ctx, request.ServerID, request.Ip, int(request.Port))
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{
		ClusterCACert: global.App.NodeInfo.ClusterCaCert,
		ClientCert:    clientCert,
		LinkStates:    convert.LinkStatesRouterToPb(router.GetLinkStates()),
	}, nil
}

// collectLatencies works in the cluster, test latency to new node, everything configured
func collectLatencies(ctx context.Context, ip string, port int, id string) (router.LinkStates, error) {
	nodes, err := NodeRepo.GetAll(ctx)
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
