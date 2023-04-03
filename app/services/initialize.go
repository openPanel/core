package services

import (
	"context"
	"net/netip"
	"sync"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/openPanel/core/app/clients/rpc"
	"github.com/openPanel/core/app/config"
	"github.com/openPanel/core/app/constant"
	. "github.com/openPanel/core/app/db/repo/shared"
	"github.com/openPanel/core/app/generated/db/shared"
	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/manager/router"
	"github.com/openPanel/core/app/tools/security"
)

var InitializeService pb.InitializeServiceServer = new(initializeService)

type initializeService struct{}

func (s *initializeService) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var authedToken string
	err := config.Load(constant.ConfigKeyAuthorizationToken, &authedToken, constant.SharedStore)
	if err != nil {
		return nil, err
	}
	if authedToken != request.Token {
		return nil, errors.New("invalid token")
	}

	clientCert, err := security.SignCsr(global.App.NodeInfo.ClusterCaCert, global.App.ClusterInfo.CaKey, request.Csr)
	if err != nil {
		return nil, err
	}

	err = NodeRepo.AddNode(ctx, request.ServerID, request.Ip, int(request.Port))
	if err != nil {
		return nil, err
	}

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

	updatedPbNode := &pb.Node{
		Id:   request.ServerID,
		Ip:   request.Ip,
		Port: int64(request.Port),
	}

	errs := make([]error, 0)
	wg := sync.WaitGroup{}
	wg.Add(len(nodes))
	for _, node := range nodes {
		if node.ID == request.ServerID {
			wg.Done()
			continue
		}
		go func(node *shared.Node) {
			defer wg.Done()
			err := rpc.NotifyNodeUpdate(node.ID, updatedPbNode)
			if err != nil {
				log.Warnf("failed to notify node %s of update: %s", node.ID, err)
				errs = append(errs, err)
			}
		}(node)
	}
	wg.Wait()

	if len(errs) > 0 {
		singleErr := errors.New("failed to notify all nodes of update")
		for _, err := range errs {
			singleErr = errors.Wrap(singleErr, err.Error())
		}
		return nil, singleErr
	}

	currentRouterInfos := router.GetRouterInfo()

	// after this op router will fall back to default algorithm
	// thus we should make new node broadcast its link state immediately, to resume to dijkstra soon
	router.AddNodes([]router.Node{
		{
			Id:       request.ServerID,
			AddrPort: netip.AddrPortFrom(netip.MustParseAddr(request.Ip), uint16(request.Port)),
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
		Nodes:         pbNodes,
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

	currentRouterInfos := router.GetRouterInfo()

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
