package quicgrpc

import (
	"context"
	"crypto/tls"
	"net"

	"google.golang.org/grpc/credentials"
)

var _ credentials.AuthInfo = (*Info)(nil)

// Info contains the auth information
type Info struct {
	conn *Conn
}

func NewInfo(conn *Conn) *Info {
	return &Info{conn}
}

func (i *Info) AuthType() string {
	return "quic-tls"
}

var _ credentials.TransportCredentials = (*Credentials)(nil)

type Credentials struct {
	tlsConfig        *tls.Config
	isQuicConnection bool
	serverName       string

	grpcCreds credentials.TransportCredentials
}

func NewCredentials(tlsConfig *tls.Config) credentials.TransportCredentials {
	grpcCreds := credentials.NewTLS(tlsConfig)
	return &Credentials{
		tlsConfig: tlsConfig,
		grpcCreds: grpcCreds,
	}
}

func (c *Credentials) ClientHandshake(ctx context.Context, s string, conn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	if co, ok := conn.(*Conn); ok {
		c.isQuicConnection = true
		return conn, NewInfo(co), nil
	}

	return c.grpcCreds.ClientHandshake(ctx, s, conn)
}

func (c *Credentials) ServerHandshake(conn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	if co, ok := conn.(*Conn); ok {
		c.isQuicConnection = true
		return conn, NewInfo(co), nil
	}

	return c.grpcCreds.ServerHandshake(conn)
}

func (c *Credentials) Info() credentials.ProtocolInfo {
	if c.isQuicConnection {
		return credentials.ProtocolInfo{
			SecurityProtocol: "quic-tls",
			ServerName:       c.serverName,
			ProtocolVersion:  "/quic/1.0.0",
		}
	}

	return c.grpcCreds.Info()
}

func (c *Credentials) Clone() credentials.TransportCredentials {
	return &Credentials{
		tlsConfig: c.tlsConfig.Clone(),
		grpcCreds: c.grpcCreds.Clone(),
	}
}

// OverrideServerName deprecated, make golang happy :(
func (c *Credentials) OverrideServerName(s string) error {
	c.serverName = s
	return c.grpcCreds.OverrideServerName(s)
}
