package rpcDialer

import (
	"google.golang.org/grpc"

	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/manager/router"
	"github.com/openPanel/core/app/tools/ca"
	"github.com/openPanel/core/app/tools/middleware/client"
	"github.com/openPanel/core/app/tools/quicNet"
)

func DialWithAddress(address, target string) (*grpc.ClientConn, error) {
	cachedConn := getClientConnFromCache(target)
	if cachedConn != nil {
		return cachedConn, nil
	}

	tlsConfig, err := ca.GenerateRPCTLSConfig(
		global.App.NodeInfo.ServerCert,
		global.App.NodeInfo.ServerPrivateKey,
		global.App.NodeInfo.ClusterCaCert,
	)
	if err != nil {
		log.Errorf("error generating tls config: %v", err)
		return nil, err
	}

	creds := quicNet.NewCredentials(tlsConfig)
	dialer := quicNet.NewQuicDialer(tlsConfig)

	options := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
		grpc.WithContextDialer(dialer),
		client.GetStreamInterceptorOption(global.App.NodeInfo.ServerId, target),
		client.GetUnaryInterceptorOption(global.App.NodeInfo.ServerId, target),
		grpc.WithAuthority(target),
	}
	conn, err := grpc.Dial(address, options...)
	if err != nil {
		log.Warnf("error dialing:%s %v", address, err)
		return nil, err
	}

	saveClientConnToCache(target, conn)

	return conn, nil
}

func DialWithServerId(serverId string) (*grpc.ClientConn, error) {
	address, err := router.GetHop(serverId)
	if err != nil {
		return nil, err
	}
	return DialWithAddress(address.String(), serverId)
}
