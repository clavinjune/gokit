package proxyutil

import (
	"net"

	"net/http"

	"github.com/clavinjune/gokit/grpcutil"
	"github.com/google/wire"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

var (
	ProviderSet = wire.NewSet(
		NewServer,
		NewServerOption,
	)
)

func NewServer(opt *ServerOption) *Server {
	return &Server{
		core: &http.Server{
			Handler: http.NotFoundHandler(),
		},
		listener: opt.Listener,
		logger: opt.Logger.With(
			slog.String("server", "PROXY"),
		),
		proxy: opt.Server,

		grpcAddress:  opt.GrpcAddress,
		grpcDialOpts: opt.GrpcDialOpts,
	}
}

func NewServerOption(
	grpcAddress grpcutil.Address, grpcDialOpts []grpc.DialOption,
	addr Address, logger *slog.Logger, opts ...runtime.ServeMuxOption) (*ServerOption, func(), error) {
	l, err := net.Listen("tcp", addr.Address())
	if err != nil {
		return nil, func() {}, err
	}

	opt := &ServerOption{
		GrpcAddress:  grpcAddress,
		GrpcDialOpts: grpcDialOpts,
		Listener:     l,
		Logger:       logger,
		Server:       runtime.NewServeMux(opts...),
	}

	cleaner := func() {
		l.Close()
	}
	return opt, cleaner, nil
}
