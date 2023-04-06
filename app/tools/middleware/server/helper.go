package server

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/openPanel/core/app/constant"
)

// getReplyInstanceFromHandler fucking ugly and slow hack, looking for a better solution
func getReplyInstanceFromHandler(fn grpc.UnaryHandler) any {
	t := reflect.TypeOf(fn)
	ret := t.Out(0)
	return reflect.New(ret).Interface()
}

func getDataInstanceFromTransferFn(fn any) any {
	t := reflect.TypeOf(fn)
	ret := t.In(0)
	return reflect.New(ret).Interface()
}

func getDstFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("No metadata found")
	}
	dsts := md.Get(constant.RPCDestinationMetadataKey)
	if dsts == nil {
		return "", errors.New("No metadata found")
	}
	if len(dsts) < 1 {
		return "", errors.New("destination not attached")
	}
	return dsts[0], nil
}
