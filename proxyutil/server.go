package proxyutil

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"

	"net/http"

	"github.com/clavinjune/gokit/grpcutil"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

type Server struct {
	_        struct{}
	core     *http.Server
	listener net.Listener
	logger   *slog.Logger
	proxy    *runtime.ServeMux

	grpcAddress  grpcutil.Address
	grpcDialOpts []grpc.DialOption
}

func (s *Server) Handle(fn func(proxy *runtime.ServeMux) http.HandlerFunc) *Server {
	s.core.Handler = fn(s.proxy)
	return s
}
func (s *Server) Register(fn func(proxy *runtime.ServeMux, endpoint string, dialOpts []grpc.DialOption)) *Server {
	fn(s.proxy, s.grpcAddress.Address(), s.grpcDialOpts)
	return s
}

func (s *Server) Start(_ context.Context) error {
	s.logger.LogAttrs(slog.LevelInfo, "starting",
		slog.String("address", s.listener.Addr().String()),
	)
	if err := s.core.Serve(s.listener); err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) {
	s.logger.LogAttrs(slog.LevelInfo, "stopped")
	s.core.Shutdown(ctx)
}

func (s *Server) Listen(ctx context.Context) {
	go func() {
		if err := s.Start(ctx); err != nil {
			s.logger.LogAttrs(slog.LevelError, err.Error())
		}
	}()

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM)
	stopSignal := <-stopCh

	s.logger.LogAttrs(slog.LevelInfo, "received stop signal",
		slog.Any("signal", stopSignal),
	)
	s.Stop(ctx)
}
