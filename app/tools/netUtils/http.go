package netUtils

import (
	"crypto/tls"
	"net/http"
)

func GetInsecureTLSHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}
