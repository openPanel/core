package outer

import (
	"io"
	"net"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

const CloudflareTraceEndpoint = "https://dash.cloudflare.com/cdn-cgi/trace"

func GetServerIP() (net.IP, error) {
	resp, err := http.Get(CloudflareTraceEndpoint)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	respText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	respTextStr := string(respText)

	lines := strings.Split(respTextStr, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "ip=") {
			return net.ParseIP(strings.TrimPrefix(line, "ip=")), nil
		}
	}
	return nil, errors.New("could not find ip in response")
}
