package utils

import (
	"testing"

	"go.uber.org/zap"

	"github.com/openPanel/core/app/global/log"
)

func mockLogger() {
	log.UpdateLogger(zap.NewNop().Sugar())
}

func Test_getPublicIPWithCloudflare(t *testing.T) {
	t.Run("cf ip get test", func(t *testing.T) {
		mockLogger()
		got, err := getPublicIPWithCloudflare()
		if err != nil {
			t.Errorf("getPublicIPWithCloudflare() error = %v", err)
			return
		}
		if got == nil {
			t.Errorf("getPublicIPWithCloudflare() got = %v", got)
		}
	})
}

func Test_getPublicIPWithIpify(t *testing.T) {
	t.Run("ipify ip get test", func(t *testing.T) {
		got, err := getPublicIPWithIpify()
		if err != nil {
			t.Errorf("getPublicIPWithIpify() error = %v", err)
			return
		}
		if got == nil {
			t.Errorf("getPublicIPWithIpify() got = %v", got)
		}
	})
}
