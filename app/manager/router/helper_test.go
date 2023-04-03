package router

import (
	"net/netip"
	"sync"

	"github.com/openPanel/core/app/global"
)

// since algorithm change global vars, it should only be tested sequentially
var testLock sync.Mutex

func setupTestData() {
	nodes = map[string]Node{
		"A": {
			Id:       "A",
			AddrPort: netip.MustParseAddrPort("127.0.0.1:8080"),
		},
		"B": {
			Id:       "B",
			AddrPort: netip.MustParseAddrPort("127.0.0.2:8081"),
		},
		"C": {
			Id:       "C",
			AddrPort: netip.MustParseAddrPort("127.0.0.3:8082"),
		},
		"D": {
			Id:       "D",
			AddrPort: netip.MustParseAddrPort("127.0.0.4:8083"),
		},
		"E": {
			Id:       "E",
			AddrPort: netip.MustParseAddrPort("127.0.0.5:8084"),
		},
	}

	routerInfos = map[Edge]int{
		Edge{From: "A", To: "B"}: 7,
		Edge{From: "A", To: "C"}: 15,
		Edge{From: "B", To: "C"}: 2,
		Edge{From: "B", To: "D"}: 9,
		Edge{From: "C", To: "B"}: 1,
		Edge{From: "C", To: "D"}: 3,
		Edge{From: "A", To: "E"}: 2,
		Edge{From: "B", To: "E"}: 3,
	}

	global.App.NodeInfo.ServerId = "A"
}
