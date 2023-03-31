package services

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/openPanel/core/app/config"
	"github.com/openPanel/core/app/constant"
	"github.com/openPanel/core/app/db/repo/shared"
	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/tools/security"
)

var InitializeService pb.InitializeServiceServer = new(initializeService)

type initializeService struct{}

func (i initializeService) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var authedToken string
	err := config.Load(constant.ConfigKeyAuthorizationToken, &authedToken, constant.SharedStore)
	if err != nil {
		return nil, err
	}
	if authedToken != request.Token {
		return nil, errors.New("invalid token")
	}

	resp := &pb.RegisterResponse{
		ClusterCACert: global.App.ClusterInfo.CaCert,
	}

	clientCert, err := security.SignCsr(global.App.ClusterInfo.CaCert, global.App.ClusterInfo.CaKey, request.Csr)
	if err != nil {
		return nil, err
	}
	resp.ClientCert = clientCert

	err = shared.NodeRepo.AddNode(ctx, request.ServerID, request.Ip, int(request.Port))
	if err != nil {
		return nil, err
	}

	nodes, err := shared.NodeRepo.GetAll(ctx)
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

	resp.Nodes = pbNodes

	return resp, nil
}

func (i initializeService) GetNodesInfo(ctx context.Context, empty *emptypb.Empty) (*pb.GetNodesInfoResponse, error) {
	nodes, err := shared.NodeRepo.GetAll(ctx)
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

	return &pb.GetNodesInfoResponse{
		Nodes: pbNodes,
	}, nil
}
