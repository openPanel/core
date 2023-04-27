package netUtils

import (
	"crypto/tls"
	"net/http"
	"time"
)

func GetInsecureHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Second * 10,
	}
}
