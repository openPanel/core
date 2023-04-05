package gateway

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/openPanel/core/app/constant"
)

var _ runtime.HeaderMatcherFunc = CustomMatcher

func CustomMatcher(key string) (string, bool) {
	switch key {
	case constant.RPCDestinationMetadataKey, constant.RPCSourceMetadataKey:
		return key, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}
