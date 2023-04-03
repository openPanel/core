package http

import (
	"fmt"
	"net/netip"
	"time"

	"github.com/openPanel/core/app/tools/utils/netUtils"
)

func QuicPing(target netip.AddrPort) (int, error) {
	client := netUtils.GetInsecureQuicHttpClient()

	start := time.Now()
	_, err := client.Head(fmt.Sprintf("https://%s/", target.String()))
	if err != nil {
		return -1, err
	}
	latency := time.Since(start).Milliseconds()
	return int(latency), nil
}
