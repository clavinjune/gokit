package grpcutil

import (
	"net"
	"net/http"

	"github.com/google/wire"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

var (
	ProviderSet = wire.NewSet(
		NewServer,
		NewServerOption,
		NewProxy,
		NewProxyOption,
	)
)

func NewServer(opt *ServerOption) *Server {
	return &Server{
		logger: opt.Logger.With(
			slog.String("server", "GRPC"),
		),
		core:     opt.Server,
		Listener: opt.Listener,
	}
}

func NewProxy(opt *ProxyOption) *Proxy {
	core := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			opt.Proxy.ServeHTTP(w, r)
		}),
	}

	return &Proxy{
		logger: opt.Logger.With(
			slog.String("server", "HTTP"),
		),
		core:       core,
		proxy:      opt.Proxy,
		Listener:   opt.Listener,
		grpcServer: opt.GrpcServer,
	}
}

func NewServerOption(port ServerPort, server *grpc.Server, logger *slog.Logger) (
	*ServerOption, func(), error,
) {
	l, err := net.Listen("tcp", port.Address())
	if err != nil {
		return nil, func() {}, err
	}

	opt := &ServerOption{
		Listener: l,
		Server:   server,
		Logger:   logger,
	}

	cleaner := func() {
		l.Close()
	}

	return opt, cleaner, nil
}

func NewProxyOption(server *Server, proxyPort ProxyPort,
	muxOpts []runtime.ServeMuxOption, logger *slog.Logger,
) (*ProxyOption, func(), error) {

	l, err := net.Listen("tcp", proxyPort.Address())
	if err != nil {
		return nil, func() {}, err
	}

	mux := runtime.NewServeMux(muxOpts...)
	opt := &ProxyOption{
		GrpcServer: server,
		Listener:   l,
		Proxy:      mux,
		Logger:     logger,
	}

	cleaner := func() {
		l.Close()
	}
	return opt, cleaner, nil
}
