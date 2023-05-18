package router

import (
	"fmt"
	"math/rand"
	"net/netip"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
)

func setupTestData(t *testing.T) {
	t.Helper()

	nodes = map[string]netip.AddrPort{
		"A": netip.MustParseAddrPort("127.0.0.1:8080"),
		"B": netip.MustParseAddrPort("127.0.0.2:8081"),
		"C": netip.MustParseAddrPort("127.0.0.3:8082"),
		"D": netip.MustParseAddrPort("127.0.0.4:8083"),
		"E": netip.MustParseAddrPort("127.0.0.5:8084"),
	}

	linkStates = map[string][]LinkState{
		"A": {
			{From: "A", To: "B", Latency: 7},
			{From: "A", To: "C", Latency: 15},
			{From: "A", To: "E", Latency: 2},
		},
		"B": {
			{From: "B", To: "C", Latency: 2},
			{From: "B", To: "D", Latency: 9},
			{From: "B", To: "E", Latency: 3},
		},
		"C": {
			{From: "C", To: "B", Latency: 1},
			{From: "C", To: "D", Latency: 3},
		},
	}

	global.App.NodeInfo.ServerId = "A"
}

func setupBenchmarkData(b *testing.B, size int) {
	b.Helper()
	defer b.ResetTimer()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	nodes = make(map[string]netip.AddrPort, size)
	linkStates = make(map[string][]LinkState, size)

	getRandomAddrPort := func() netip.AddrPort {
		return netip.MustParseAddrPort(fmt.Sprintf("%d.%d.%d.%d:%d", r.Intn(255), r.Intn(255), r.Intn(255), r.Intn(255), r.Intn(65535)))
	}

	keys := make([]string, 0, size)

	for i := 0; i < size; i++ {
		keys = append(keys, uuid.New().String())
		nodes[keys[i]] = getRandomAddrPort()
	}

	for node := range nodes {
		from := node

		tos := make(map[string]struct{}, size/3)

		for i := 0; i < size/3; i++ {
			to := keys[r.Intn(size)]

			if _, ok := tos[to]; ok {
				i--
				continue
			}

			tos[to] = struct{}{}

			linkStates[from] = append(linkStates[from], LinkState{
				From:    from,
				To:      to,
				Latency: r.Intn(100),
			})
		}
	}

	global.App.NodeInfo.ServerId = keys[rand.Intn(size)]

	log.Warnf = func(template string, args ...any) {
		println(fmt.Sprintf(template, args...))
	}
}
