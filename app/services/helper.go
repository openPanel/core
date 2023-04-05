package services

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/openPanel/core/app/constant"
	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/tools/middleware/server"
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
	srcs := md.Get(constant.RPCSourceMetadataKey)
	if srcs == nil || len(srcs) < 1 {
		return "", "", errors.New("No source attached")
	}
	src := srcs[0]
	return src, dst, nil
}

func newGrpcServer() *grpc.Server {
	grpcServer := grpc.NewServer(server.GetServerUnaryInterceptorOption(), server.GetServerStreamInterceptorOption())

	registerServers(grpcServer)

	return grpcServer
}

func registerServers(server *grpc.Server) {
	pb.RegisterLinkStateServiceServer(server, LinkStateService)
	pb.RegisterInitializeServiceServer(server, InitializeService)
	pb.RegisterDqliteConnectionServer(server, DqliteService)
}
