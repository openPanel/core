package rpc

import (
	"context"
	"errors"
	"sync"

	"github.com/openPanel/core/app/db/repo/shared"
	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/tools/rpcDialer"
)

type BroadcastMessage = struct {
	Type    pb.BroadcastType
	Payload string
}

func Broadcast(messages []BroadcastMessage) error {
	nodes, err := shared.NodeRepo.GetBroadcastNodes(context.Background())
	if err != nil {
		return err
	}

	request := &pb.MultiBroadcastRequest{
		Broadcasts: make([]*pb.Broadcast, len(messages)),
	}

	usedType := make(map[pb.BroadcastType]bool)

	for i, message := range messages {
		if _, ok := usedType[message.Type]; ok {
			return errors.New("duplicate broadcast type")
		}
		request.Broadcasts[i] = &pb.Broadcast{
			Type:    message.Type,
			Payload: message.Payload,
		}
		usedType[message.Type] = true
	}

	wg := sync.WaitGroup{}
	wg.Add(len(nodes))

	errs := make([]error, 0)

	for i := range nodes {
		go func(i int) {
			defer wg.Done()

			node := nodes[i]
			conn, err := rpcDialer.DialWithServerId(node.ID)
			if err != nil {
				errs = append(errs, err)
				return
			}

			client := pb.NewBroadcastServiceClient(conn)
			_, err = client.Broadcast(context.Background(), request)

			if err != nil {
				errs = append(errs, err)
			}
		}(i)
	}

	wg.Wait()
	if len(errs) > 0 {
		err := errors.Join(errs...)
		return err
	}

	return nil
}
