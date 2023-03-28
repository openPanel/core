package http

import (
	"time"

	"github.com/openPanel/core/app/tools/utils/netUtils"
)

func QuicPing(url string) (int, error) {
	client := netUtils.GetInsecureQuicHttpClient()

	start := time.Now()
	_, err := client.Head(url)
	if err != nil {
		return -1, err
	}
	latency := time.Since(start).Milliseconds()
	return int(latency), nil
}
