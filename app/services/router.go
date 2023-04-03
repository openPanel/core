package services

import (
	"context"
	"net"
	"net/netip"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/openPanel/core/app/config"
	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/manager/router"
)

var LinkStateService pb.LinkStateServiceServer = new(linkStateService)

type linkStateService struct{}

func (l linkStateService) UpdateLinkState(ctx context.Context, request *pb.LinkStateUpdateRequest) (*emptypb.Empty, error) {
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

func (l linkStateService) NotifyNodeUpdate(ctx context.Context, request *pb.NodeUpdateRequest) (*emptypb.Empty, error) {
	ip := net.ParseIP(request.UpdatedNode.Ip)
	if ip == nil {
		return nil, errors.New("invalid ip")
	}

	id := request.UpdatedNode.Id
	port := int(request.UpdatedNode.Port)

	router.UpdateNode(id, ip, port)

	err := config.AppendNodesCache(config.NodeCacheEntry{
		Id: id,
		AddrPort: netip.AddrPortFrom(
			netip.MustParseAddr(ip.String()),
			uint16(port)),
	})
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
