package global

import (
	"net"
)

type NodeInfo struct {
	ServerId         string
	ServerPublicIP   net.IP
	ServerListenIP   net.IP
	ServerPort       int
	ServerCert       []byte
	ServerPrivateKey []byte

	ClusterCaCert []byte

	IsIndirectIP bool
}

type ClusterInfo struct {
	CaKey []byte
}
