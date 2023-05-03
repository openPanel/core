package services

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/tools/broadcast"
)

var BroadcastService pb.BroadcastServiceServer = new(broadcastService)

type broadcastService struct{}

func (b broadcastService) Broadcast(ctx context.Context, request *pb.MultiBroadcastRequest) (*emptypb.Empty, error) {
	usedType := make(map[pb.BroadcastType]bool)
	for _, bc := range request.Broadcasts {
		if usedType[bc.Type] {
			return nil, errors.New("duplicate broadcast type")
		}
		usedType[bc.Type] = true
	}

	errs := make([]error, 0)
	for _, bc := range request.Broadcasts {
		handler := broadcastHandler[bc.Type]
		if handler == nil {
			errs = append(errs, errors.New("unknown broadcast type"))
			continue
		}

		err := handler(ctx, bc.Payload)
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

type BroadcastHandler func(ctx context.Context, payload []byte) error

var broadcastHandler = map[pb.BroadcastType]BroadcastHandler{
	pb.BroadcastType_LINK_STATE_CHANGE: linkStateChange,
}

var _ BroadcastHandler = linkStateChange

func linkStateChange(_ context.Context, payload []byte) error {
	return broadcast.LoadRouterPayload(payload)
}
