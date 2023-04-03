package bootstrap

import (
	"net"
	"net/netip"
	"strconv"

	"github.com/openPanel/core/app/generated/db/shared"
	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/manager/dqlite"
)

func createDqlite() *shared.Client {
	return dqlite.CreateSharedDatabase(nil)
}

func joinDqlite(info *pb.RegisterResponse) *shared.Client {
	addrs := make([]string, len(info.Nodes))
	for i, node := range info.Nodes {
		addrs[i] = net.JoinHostPort(node.Ip, strconv.Itoa(int(node.Port)))
	}
	return dqlite.CreateSharedDatabase(&addrs)
}

func resumeDqlite(addrs []netip.AddrPort) *shared.Client {
	if len(addrs) == 0 {
		return dqlite.CreateSharedDatabase(nil)
	}

	addresses := make([]string, len(addrs))
	for i, addr := range addrs {
		addresses[i] = addr.String()
	}
	return dqlite.CreateSharedDatabase(&addresses)
}
