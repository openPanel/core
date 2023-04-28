package services

import (
	"context"
	"encoding/json"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/openPanel/core/app/clients/rpc"
	"github.com/openPanel/core/app/clients/tcp"
	. "github.com/openPanel/core/app/db/repo/shared"
	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/manager/router"
	"github.com/openPanel/core/app/tools/ca"
	"github.com/openPanel/core/app/tools/convert"
	"github.com/openPanel/core/app/tools/utils/netUtils"
)

var InitializeService pb.InitializeServiceServer = new(initializeService)

type initializeService struct{}

func (s *initializeService) UpdateLinkState(ctx context.Context, request *pb.UpdateLinkStateRequest) (*pb.UpdateLinkStateResponse, error) {
	lstFromNewNode := convert.LinkStatesPbToRouter(request.LinkStates)
	lstToNewNode, err := rpc.CollectLatencies(ctx, request.Ip, int(request.Port), request.ServerID)
	if err != nil {
		return nil, err
	}

	router.UpdateLinkStates(lstFromNewNode, lstToNewNode)

	fullLst := router.GetLinkStates()

	broadcastPayload, err := json.Marshal(router.GetLinkStates())
	if err != nil {
		return nil, err
	}

	// broadcast new link state
	err = rpc.Broadcast([]rpc.BroadcastMessage{{
		Type:    pb.BroadcastType_NOTIFY_LINK_STATE_CHANGE,
		Payload: string(broadcastPayload),
	}})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateLinkStateResponse{
		LinkStates: convert.LinkStatesRouterToPb(fullLst),
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

	err = NodeRepo.AddNode(ctx, request.ServerID, request.Ip, int(request.Port))
	if err != nil {
		return nil, err
	}

	lstFromNewNode := convert.LinkStatesPbToRouter(request.LinkStates)
	lstToNewNode, err := rpc.CollectLatencies(ctx, request.Ip, int(request.Port), request.ServerID)
	if err != nil {
		return nil, err
	}

	router.UpdateLinkStates(lstFromNewNode, lstToNewNode)

	fullLst := router.GetLinkStates()

	broadcastPayload, err := json.Marshal(router.GetLinkStates())
	if err != nil {
		log.Errorf("failed to marshal link states: %v", err)
		return nil, err
	}

	err = rpc.Broadcast([]rpc.BroadcastMessage{{
		Type: pb.BroadcastType_NOTIFY_NODE_CHANGE,
	}, {
		Type:    pb.BroadcastType_NOTIFY_LINK_STATE_CHANGE,
		Payload: string(broadcastPayload),
	}})
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{
		ClusterCACert: global.App.NodeInfo.ClusterCaCert,
		ClientCert:    clientCert,
		LinkStates:    convert.LinkStatesRouterToPb(fullLst),
	}, nil
}
