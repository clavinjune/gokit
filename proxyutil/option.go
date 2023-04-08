package proxyutil

import (
	"net"

	"github.com/clavinjune/gokit/grpcutil"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

type ServerOption struct {
	_            struct{}
	GrpcAddress  grpcutil.Address
	GrpcDialOpts []grpc.DialOption
	Listener     net.Listener
	Logger       *slog.Logger
	Server       *runtime.ServeMux
}
