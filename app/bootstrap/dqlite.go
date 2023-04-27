package bootstrap

import (
	"net/netip"

	"github.com/openPanel/core/app/generated/db/shared"
	"github.com/openPanel/core/app/manager/dqlite"
	"github.com/openPanel/core/app/manager/router"
)

// for initial node
func createDqlite() *shared.Client {
	return dqlite.CreateSharedDatabase(nil)
}

// for success join, len(nodes) >= 1
func joinDqlite(nodes []router.Node) *shared.Client {
	addrs := make([]string, len(nodes))
	for i, node := range nodes {
		addrs[i] = node.AddrPort.String()
	}
	return dqlite.CreateSharedDatabase(&addrs)
}

// at resume, len(addrs) >= 0, if len(addrs) == 0, it means that the node is the first node
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
