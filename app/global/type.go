package global

import (
	"net"
)

type NodeInfo struct {
	ServerId         string
	ServerIp         net.IP
	ServerPort       int
	ServerCert       []byte
	ServerPrivateKey []byte

	IsIndirectIP bool
}

type ClusterInfo struct {
	CaCert []byte
	CaKey  []byte
}
