package constant

import (
	"net"
)

var DefaultListenIp = net.IPv4zero

const DefaultListenPort = 40123
const DefaultDataDir = "data"

const DefaultLocalSqliteFilename = "core.local.db"

const RpcUnixListenAddress = "/tmp/openPanel.sock"
const RpcUnixDialAddress = "unix-abstract:/tmp/openPanel.sock"
