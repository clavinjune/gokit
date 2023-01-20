package grpcutil

import (
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

type ServerOption struct {
	_        struct{}
	Listener net.Listener
	Server   *grpc.Server
	Logger   *slog.Logger
}

type ProxyOption struct {
	_              struct{}
	GrpcServer     *Server
	Listener       net.Listener
	Proxy          *runtime.ServeMux
	Logger         *slog.Logger
	OpenApiHandler http.Handler
}
