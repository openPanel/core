package middleware

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/providers/zap/v2"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"

	"github.com/openPanel/core/app/global"
)

var ServerUnaryInterceptorOption = grpc.ChainUnaryInterceptor(
	logging.UnaryServerInterceptor(zap.InterceptorLogger(global.App.Logger.Named("grpc"))),
)

var ServerStreamInterceptorOption = grpc.ChainStreamInterceptor(
	logging.StreamServerInterceptor(zap.InterceptorLogger(global.App.Logger.Named("grpc"))),
)

var ClientUnaryInterceptorOption = grpc.WithChainUnaryInterceptor(
	logging.UnaryClientInterceptor(zap.InterceptorLogger(global.App.Logger.Named("grpc"))),
)

var ClientStreamInterceptorOption = grpc.WithChainStreamInterceptor(
	logging.StreamClientInterceptor(zap.InterceptorLogger(global.App.Logger.Named("grpc"))),
)
