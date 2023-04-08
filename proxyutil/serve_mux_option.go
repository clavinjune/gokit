package proxyutil

import "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

func RemoveOutgoingHeader(whitelist map[string]string) runtime.ServeMuxOption {
	return runtime.WithOutgoingHeaderMatcher(func(s string) (string, bool) {
		val, ok := whitelist[s]
		return val, ok
	})
}
