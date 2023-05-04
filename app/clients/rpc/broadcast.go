package rpc

import (
	"context"
	"sync"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/openPanel/core/app/db/repo/shared"
	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/tools/rpcDialer"
)

type BroadcastMessage = struct {
	Type    pb.BroadcastType
	Payload []byte
}

// Broadcast used for single shot broadcast
func Broadcast(message BroadcastMessage) error {
	nodes, err := shared.NodeRepo.GetBroadcastNodes(context.Background())
	if err != nil {
		return err
	}

	request := &pb.MultiBroadcastRequest{
		Broadcasts: []*pb.Broadcast{{
			Type:    message.Type,
			Payload: message.Payload,
		}},
	}

	wg := sync.WaitGroup{}
	wg.Add(len(nodes))

	errChan := make(chan error, len(nodes))
	defer close(errChan)

	for _, node := range nodes {
		go func(id string) {
			defer wg.Done()

			conn, err := rpcDialer.DialWithServerId(id)
			if err != nil {
				errChan <- err
				return
			}
			defer func(conn *grpc.ClientConn) {
				_ = conn.Close()
			}(conn)

			client := pb.NewBroadcastServiceClient(conn)
			_, err = client.Broadcast(context.Background(), request)

			if err != nil {
				errChan <- err
				return
			}
		}(node.ID)
	}

	wg.Wait()

	if len(errChan) > 0 {
		err := errors.New("broadcast failed")
		for e := range errChan {
			err = errors.Wrap(err, e.Error())
		}
		return err
	}

	return nil
}
