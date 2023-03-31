package services

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/quic-go/quic-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

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

	tlsConfig, err := security.GenerateRPCTLSConfig(global.App.NodeInfo.ServerCert, global.App.NodeInfo.ServerPrivateKey, global.App.ClusterInfo.CaCert)
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

func StartHttpServiceBlocking() {
	grpcMux := runtime.NewServeMux()

	err := pb.RegisterInitializeServiceHandlerFromEndpoint(
		context.Background(),
		grpcMux,
		constant.RpcUnixListenAddress,
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		log.Fatalf("error registering grpc gateway: %v", err)
		return
	}

	unixListener, err := net.Listen("unix", "")
	if err != nil {
		log.Fatalf("error listening: %v", err)
	}

	go func() {
		if err := grpcServer.Serve(unixListener); err != nil {
			log.Fatalf("error serving grpc: %v", err)
		}
	}()

	tlsConfig, err := security.GenerateHTTPTLSConfig(global.App.NodeInfo.ServerCert, global.App.NodeInfo.ServerPrivateKey)
	if err != nil {
		log.Fatalf("error generating tls config: %v", err)
	}

	var addr string
	if global.App.NodeInfo.IsIndirectIP {
		addr = constant.DefaultListenIp.String()
	} else {
		addr = global.App.NodeInfo.ServerIp.String()
	}

	s := &http.Server{
		Addr:      fmt.Sprintf("%s:%d", addr, global.App.NodeInfo.ServerPort),
		TLSConfig: tlsConfig,
		Handler:   grpcMux,
	}

	if err = s.ListenAndServeTLS("", ""); err != nil {
		log.Fatalf("error serving http: %v", err)
	}
}
