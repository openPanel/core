package bootstrap

import (
	"log"
	"net"
	"os"
	"os/user"

	"github.com/openPanel/core/app/config"
	"github.com/openPanel/core/app/constant"
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

func saveNodeInfo(id string, ip net.IP, port int, cert, key []byte, indirect bool) (global.NodeInfo, error) {
	i := global.NodeInfo{
		ServerCert:       cert,
		ServerId:         id,
		ServerIp:         ip,
		ServerPort:       port,
		ServerPrivateKey: key,
		IsIndirectIP:     indirect,
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

func requireInitialStartUp() {
	// check file exist
	if _, err := os.Stat(constant.DefaultDataDir + string(os.PathSeparator) + constant.DefaultLocalSqliteFilename); err == nil {
		log.Fatal("This program has already been initialized, boot with no command to resume the server")
	} else if !os.IsNotExist(err) {
		log.Fatalf("Failed to check if database file exists: %v", err)
	}
}
