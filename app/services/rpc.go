package services

import (
	"fmt"

	"github.com/quic-go/quic-go"
	"google.golang.org/grpc"

	"github.com/openPanel/core/app/constant"
	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/tools/middleware/server"
	"github.com/openPanel/core/app/tools/quicNet"
	"github.com/openPanel/core/app/tools/security"
)

var grpcServer *grpc.Server

func StartRpcServiceBlocking() {
	grpcServer = grpc.NewServer(server.GetServerUnaryInterceptorOption(), server.GetServerStreamInterceptorOption())

	pb.RegisterLinkStateServiceServer(grpcServer, LinkStateService)
	pb.RegisterInitializeServiceServer(grpcServer, InitializeService)
	pb.RegisterDqliteConnectionServer(grpcServer, DqliteService)

	tlsConfig, err := security.GenerateRPCTLSConfig(global.App.NodeInfo.ServerCert, global.App.NodeInfo.ServerPrivateKey, global.App.NodeInfo.ClusterCaCert)
	if err != nil {
		log.Fatalf("error generating tls config: %v", err)
	}

	var listenAddr string

	if global.App.NodeInfo.IsIndirectIP {
		listenAddr = fmt.Sprintf("%s:%d", constant.DefaultListenIp, global.App.NodeInfo.ServerPort)
	} else {
		listenAddr = fmt.Sprintf("%s:%d", global.App.NodeInfo.ServerIp.String(), global.App.NodeInfo.ServerPort)
	}

	ql, err := quic.ListenAddr(listenAddr, tlsConfig, nil)
	if err != nil {
		log.Fatalf("error listening: %v", err)
	}
	listener := quicNet.Listen(ql)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("error serving grpc: %v", err)
	}
}
