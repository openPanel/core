package services

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"

	"github.com/openPanel/core/app/constant"
)

func getSrcAndDstFromContext(ctx context.Context) (string, string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", "", errors.New("No metadata attached")
	}

	dsts := md.Get(constant.RPCDestinationMetadataKey)
	if dsts == nil || len(dsts) < 1 {
		return "", "", errors.New("No destination attached")
	}
	dst := dsts[0]
	srcs, ok := md[constant.RPCSourceMetadataKey]
	if !ok || len(srcs) < 1 {
		return "", "", errors.New("No source attached")
	}
	src := srcs[0]
	return src, dst, nil
}
