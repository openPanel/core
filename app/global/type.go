package global

import (
	"net"
)

type LocalNodeInfo struct {
	ServerId         string
	ServerIp         net.IP
	ServerPort       int
	ServerCert       []byte
	ServerPrivateKey []byte
}

type ClusterInfo struct {
	CaCert []byte
	CaKey  []byte
}
