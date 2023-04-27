package global

import (
	"net"
)

type NodeInfo struct {
	ServerId         string
	ServerPublicIP   net.IP
	ServerListenIP   net.IP
	IsIndirectIP     bool
	ServerPort       int
	ServerCert       []byte
	ServerPrivateKey []byte

	ClusterCaCert []byte
}

type ClusterInfo struct {
	CaKey []byte
}
