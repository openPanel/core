package services

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/openPanel/core/app/generated/pb"
)

var BroadcastService pb.BroadcastServiceServer = new(broadcastService)

type broadcastService struct{}

func (b broadcastService) SingleBroadcast(ctx context.Context, request *pb.SingleBroadcastRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (b broadcastService) MultiBroadcast(ctx context.Context, request *pb.MultiBroadcastRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

var broadcastHandler = map[pb.BroadcastType]func(ctx context.Context){
	pb.BroadcastType_NOTIFY_LINK_STATE_CHANGE: linkStateChange,
	pb.BroadcastType_NOTIFY_NODE_CHANGE: nodeChange,
}

func linkStateChange(ctx context.Context) {

}

func nodeChange(ctx context.Context) {

}
