package services

import (
	"fmt"

	"github.com/quic-go/quic-go"

	"github.com/openPanel/core/app/bootstrap/clean"
	"github.com/openPanel/core/app/constant"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/tools/quicNet"
	"github.com/openPanel/core/app/tools/security"
)

func StartRpcServiceBlocking() {
	grpcServer := newGrpcServer()

	tlsConfig, err := security.GenerateRPCTLSConfig(global.App.NodeInfo.ServerCert, global.App.NodeInfo.ServerPrivateKey, global.App.NodeInfo.ClusterCaCert)
	if err != nil {
		log.Fatalf("error generating tls config: %v", err)
	}

	var listenAddr string

	if global.App.NodeInfo.IsIndirectIP {
		listenAddr = fmt.Sprintf("%s:%d", constant.DefaultListenIp, global.App.NodeInfo.ServerPort)
	} else {
		listenAddr = fmt.Sprintf("%s:%d", global.App.NodeInfo.ServerListenIP.String(), global.App.NodeInfo.ServerPort)
	}

	ql, err := quic.ListenAddr(listenAddr, tlsConfig, nil)
	if err != nil {
		log.Fatalf("error listening: %v", err)
	}
	listener := quicNet.Listen(ql)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("error serving grpc: %v", err)
	}

	clean.RegisterCleanup(func() {
		grpcServer.GracefulStop()
		log.Debug("grpc service stopped")
	})
}
