package router

import (
	"net"
	"sync"

	"github.com/openPanel/core/app/global"
)

// since algorithm change global vars, it should only be tested sequentially
var testLock sync.Mutex

func setupTestData() {
	nodes = map[string]Node{
		"A": {
			Id: "A",
			Address: Address{
				Ip:   net.IPv4(127, 0, 0, 1),
				Port: 8080,
			},
		},
		"B": {
			Id: "B",
			Address: Address{
				Ip:   net.IPv4(127, 0, 0, 2),
				Port: 8081,
			},
		},
		"C": {
			Id: "C",
			Address: Address{
				Ip:   net.IPv4(127, 0, 0, 3),
				Port: 8082,
			},
		},
		"D": {
			Id: "D",
			Address: Address{
				Ip:   net.IPv4(127, 0, 0, 4),
				Port: 8083,
			},
		},
		"E": {
			Id: "E",
			Address: Address{
				Ip:   net.IPv4(127, 0, 0, 5),
				Port: 8084,
			},
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
