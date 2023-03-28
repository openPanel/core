package bootstrap

import (
	"log"
	"net"
	"os/user"

	"github.com/openPanel/core/app/config"
	"github.com/openPanel/core/app/global"

	userLog "github.com/openPanel/core/app/global/log"
)

func init() {
	RequireRoot()
}

func RequireRoot() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	if u.Uid != "0" {
		log.Fatal("This program requires root privileges")
	}
}

func saveNodeInfo(id string, ip net.IP, port int, cert, key []byte) (global.LocalNodeInfo, error) {
	i := global.LocalNodeInfo{
		ServerCert:       cert,
		ServerId:         id,
		ServerIp:         ip,
		ServerPort:       port,
		ServerPrivateKey: key,
	}
	err := config.SaveLocalNodeInfo(i)
	return i, err
}

func loadNodeInfo() (id string, ip net.IP, port int, cert, key []byte) {
	i, err := config.LoadLocalNodeInfo()
	if err != nil {
		userLog.Fatalf("Failed to load node info: %v", err)
	}
	return i.ServerId, i.ServerIp, i.ServerPort, i.ServerCert, i.ServerPrivateKey
}
