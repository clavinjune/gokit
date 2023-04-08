package grpcutil

import (
	"net"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

type ServerOption struct {
	_        struct{}
	Listener net.Listener
	Logger   *slog.Logger
	Server   *grpc.Server
}
