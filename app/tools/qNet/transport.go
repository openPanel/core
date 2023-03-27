package qNet

import (
	"context"
	"crypto/tls"
	"net"

	"github.com/quic-go/quic-go"
	"google.golang.org/grpc/credentials"
)

var _ credentials.AuthInfo = (*Info)(nil)

// Info contains the auth information
type Info struct {
	quic.Connection
}

func NewInfo(conn quic.Connection) *Info {
	return &Info{conn}
}

func (i *Info) AuthType() string {
	return "qtls"
}

var _ credentials.TransportCredentials = (*Credentials)(nil)

type Credentials struct {
	tlsConfig        *tls.Config
	isQuicConnection bool
	serverName       string

	grpcCreds credentials.TransportCredentials
}

func NewCredentials(tlsConfig *tls.Config) *Credentials {
	grpcCreds := credentials.NewTLS(tlsConfig)
	return &Credentials{
		tlsConfig: tlsConfig,
		grpcCreds: grpcCreds,
	}
}

func (c Credentials) ClientHandshake(ctx context.Context, s string, conn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	if co, ok := conn.(*Conn); ok {
		c.isQuicConnection = true
		return conn, NewInfo(co.conn), nil
	}

	return c.grpcCreds.ClientHandshake(ctx, s, conn)
}

func (c Credentials) ServerHandshake(conn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	if co, ok := conn.(*Conn); ok {
		c.isQuicConnection = true
		return conn, NewInfo(co.conn), nil
	}

	return c.grpcCreds.ServerHandshake(conn)
}

func (c Credentials) Info() credentials.ProtocolInfo {
	if c.isQuicConnection {
		return credentials.ProtocolInfo{
			SecurityProtocol: "qtls",
			ServerName:       c.serverName,
			ProtocolVersion:  "/qtls/0.0.1",
		}
	}

	return c.grpcCreds.Info()
}

func (c Credentials) Clone() credentials.TransportCredentials {
	return &Credentials{
		tlsConfig:        c.tlsConfig.Clone(),
		isQuicConnection: c.isQuicConnection,
		serverName:       c.serverName,
		grpcCreds:        c.grpcCreds.Clone(),
	}
}

// OverrideServerName deprecated, make golang happy :(
func (c Credentials) OverrideServerName(s string) error {
	c.serverName = s
	return c.grpcCreds.OverrideServerName(s)
}
