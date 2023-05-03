package client

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/openPanel/core/app/constant"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/tools/middleware/common"
)

func GetUnaryInterceptorOption(src, dst string) grpc.DialOption {
	return grpc.WithChainUnaryInterceptor(
		getRouterInfoUnaryInterceptor(src, dst),
		logging.UnaryClientInterceptor(common.InterceptorLogger(global.App.Logger.Named("grpc-client"))),
	)
}

func GetStreamInterceptorOption(src, dst string) grpc.DialOption {
	return grpc.WithChainStreamInterceptor(
		getRouterInfoStreamInterceptor(src, dst),
		logging.StreamClientInterceptor(common.InterceptorLogger(global.App.Logger.Named("grpc-client-stream"))),
	)
}

func getRouterInfoUnaryInterceptor(src, dst string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		newCtx := metadata.AppendToOutgoingContext(ctx, constant.RPCSourceMetadataKey, src, constant.RPCDestinationMetadataKey, dst)
		return invoker(newCtx, method, req, reply, cc, opts...)
	}
}

func getRouterInfoStreamInterceptor(src, dst string) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		newCtx := metadata.AppendToOutgoingContext(ctx, constant.RPCSourceMetadataKey, src, constant.RPCDestinationMetadataKey, dst)
		return streamer(newCtx, desc, cc, method, opts...)
	}
}
