package services

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/openPanel/core/app/clients/rpc"
	. "github.com/openPanel/core/app/db/repo/shared"
	"github.com/openPanel/core/app/generated/db/shared"
	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/manager/router"
	"github.com/openPanel/core/app/tools/security"
	"github.com/openPanel/core/app/tools/utils/netUtils"
)

var InitializeService pb.InitializeServiceServer = new(initializeService)

type initializeService struct{}

func (s *initializeService) UpdateLinkState(ctx context.Context, request *pb.UpdateLinkStateRequest) (*pb.UpdateLinkStateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *initializeService) EstimateLatency(ctx context.Context, request *pb.EstimateLatencyRequest) (*pb.EstimateLatencyResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *initializeService) GetClusterInfo(ctx context.Context, empty *emptypb.Empty) (*pb.GetClusterInfoResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *initializeService) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	clientCert, err := security.SignCsr(global.App.NodeInfo.ClusterCaCert, global.App.ClusterInfo.CaKey, request.Csr)
	if err != nil {
		return nil, err
	}

	err = NodeRepo.AddNode(ctx, request.ServerID, request.Ip, int(request.Port))
	if err != nil {
		return nil, err
	}

	nodes, err := NodeRepo.GetBroadcastNodes(ctx)
	if err != nil {
		return nil, err
	}

	pbNodes := make([]*pb.Node, len(nodes))
	for i, node := range nodes {
		pbNodes[i] = &pb.Node{
			Id:   node.ID,
			Ip:   node.IP,
			Port: int64(node.Port),
		}
	}
	go func() {
		for _, node := range nodes {
			go func(node *shared.Node) {
				if node.ID == request.ServerID {
					return
				}
				err := rpc.NotifyNodeUpdate(node.ID)
				if err != nil {
					log.Warnf("failed to notify node %s of update: %s", node.ID, err)
				}
			}(node)
		}
	}()

	currentRouterInfos := router.GetLinkStates()

	// after this op router will fall back to default algorithm
	// thus we should make new node broadcast its link state immediately, to resume to dijkstra soon
	router.AddNodes([]router.Node{
		{
			Id:       request.ServerID,
			AddrPort: netUtils.NewAddrPortWithString(request.Ip, int(request.Port)),
		},
	})

	linkStates := make([]*pb.LinkState, len(currentRouterInfos))
	i := 0
	for edge, latency := range currentRouterInfos {
		linkStates[i] = &pb.LinkState{
			From:    edge.From,
			To:      edge.To,
			Latency: int32(latency),
		}
	}

	return &pb.RegisterResponse{
		ClusterCACert: global.App.NodeInfo.ClusterCaCert,
		ClientCert:    clientCert,
		LinkStates:    linkStates,
	}, nil
}

func (s *initializeService) GetNodesInfo(ctx context.Context, empty *emptypb.Empty) (*pb.GetNodesInfoResponse, error) {
	nodes, err := NodeRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	pbNodes := make([]*pb.Node, len(nodes))
	for i, node := range nodes {
		pbNodes[i] = &pb.Node{
			Id:   node.ID,
			Ip:   node.IP,
			Port: int64(node.Port),
		}
	}

	currentRouterInfos := router.GetLinkStates()

	linkStates := make([]*pb.LinkState, len(currentRouterInfos))
	i := 0
	for edge, latency := range currentRouterInfos {
		linkStates[i] = &pb.LinkState{
			From:    edge.From,
			To:      edge.To,
			Latency: int32(latency),
		}
	}

	return &pb.GetNodesInfoResponse{
		Nodes:      pbNodes,
		LinkStates: linkStates,
	}, nil
}
