package grpcutil

import (
	"net"

	"github.com/google/wire"
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
		core:     opt.Server,
		listener: opt.Listener,
		logger: opt.Logger.With(
			slog.String("server", "GRPC"),
		),
	}
}

func NewServerOption(addr Address, logger *slog.Logger, opts ...grpc.ServerOption) (
	*ServerOption, func(), error,
) {
	l, err := net.Listen("tcp", addr.Address())
	if err != nil {
		return nil, func() {}, err
	}

	opt := &ServerOption{
		Listener: l,
		Logger:   logger,
		Server:   grpc.NewServer(opts...),
	}

	cleaner := func() {
		l.Close()
	}

	return opt, cleaner, nil
}
