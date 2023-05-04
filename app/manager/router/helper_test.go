package router

import (
	"net/netip"
	"sync"
	"testing"

	"github.com/openPanel/core/app/global"
)

// since algorithm change global vars, it should only be tested sequentially
var testLock sync.Mutex

func setupTestData(t *testing.T) {
	t.Helper()

	nodes = map[string]netip.AddrPort{
		"A": netip.MustParseAddrPort("127.0.0.1:8080"),
		"B": netip.MustParseAddrPort("127.0.0.2:8081"),
		"C": netip.MustParseAddrPort("127.0.0.3:8082"),
		"D": netip.MustParseAddrPort("127.0.0.4:8083"),
		"E": netip.MustParseAddrPort("127.0.0.5:8084"),
	}

	linkStates = map[Edge]int{
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
