package services

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/manager/router"
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
}

var _ BroadcastHandler = linkStateChange

func linkStateChange(_ context.Context, payload string) error {
	return router.LoadBroadcastPayload([]byte(payload))
}
