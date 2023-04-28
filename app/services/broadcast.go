package services

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/openPanel/core/app/db/repo/shared"
	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/manager/router"
	"github.com/openPanel/core/app/tools/utils/netUtils"
)

var BroadcastService pb.BroadcastServiceServer = new(broadcastService)

type broadcastService struct{}

func (b broadcastService) Broadcast(ctx context.Context, request *pb.MultiBroadcastRequest) (*emptypb.Empty, error) {
	usedType := make(map[pb.BroadcastType]bool)
	for _, broadcast := range request.Broadcasts {
		if usedType[broadcast.Type] {
			return nil, errors.New("duplicate broadcast type")
		}
		usedType[broadcast.Type] = true
	}

	errs := make([]error, 0)
	for _, broadcast := range request.Broadcasts {
		handler := broadcastHandler[broadcast.Type]
		if handler == nil {
			errs = append(errs, errors.New("unknown broadcast type"))
			continue
		}

		err := handler(ctx, broadcast.Payload)
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		err := errors.New("broadcast failed")
		for _, e := range errs {
			err = errors.Wrap(err, e.Error())
		}
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

type BroadcastHandler func(ctx context.Context, payload string) error

var broadcastHandler = map[pb.BroadcastType]BroadcastHandler{
	pb.BroadcastType_NOTIFY_LINK_STATE_CHANGE: linkStateChange,
	pb.BroadcastType_NOTIFY_NODE_CHANGE:       nodeChange,
}

var _ BroadcastHandler = linkStateChange
var _ BroadcastHandler = nodeChange

func linkStateChange(_ context.Context, payload string) error {
	var lst = new(router.LinkStates)
	err := json.Unmarshal([]byte(payload), lst)
	if err != nil {
		return errors.Wrap(err, "unmarshal link state failed")
	}
	router.UpdateLinkStates(*lst)
	log.Debugf("link state change: %s", payload)
	return nil
}

func nodeChange(ctx context.Context, _ string) error {
	nodes, err := shared.NodeRepo.GetAll(ctx)
	if err != nil {
		return errors.Wrapf(err, "get all nodes failed")
	}

	routerNodes := make([]router.Node, len(nodes))
	for i, node := range nodes {
		routerNodes[i] = router.Node{
			Id:       node.ID,
			AddrPort: netUtils.NewAddrPortWithString(node.IP, node.Port),
		}
	}
	router.SetNodes(routerNodes)
	log.Debugf("node change: %s", nodes)
	return nil
}
