package rpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/tools/rpc"
)

// NotifyNodeUpdate part of the broadcast after a node status change
func NotifyNodeUpdate(target string, updateNode *pb.Node) error {
	conn, err := rpc.DialWithServerId(target)
	if err != nil {
		return err
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	client := pb.NewLinkStateServiceClient(conn)
	_, err = client.NotifyNodeUpdate(context.Background(), &pb.NodeUpdateRequest{
		UpdatedNode: updateNode,
	})
	return err
}

// UpdateLinkState part of the broadcast after latency estimation
func UpdateLinkState(target string, linkState *[]*pb.LinkState) error {
	conn, err := rpc.DialWithServerId(target)
	if err != nil {
		return err
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	client := pb.NewLinkStateServiceClient(conn)
	_, err = client.UpdateLinkState(context.Background(), &pb.LinkStateUpdateRequest{
		LinkState: *linkState,
	})
	return err
}
