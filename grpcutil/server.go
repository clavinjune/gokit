package grpcutil

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

type Server struct {
	_        struct{}
	core     *grpc.Server
	listener net.Listener
	logger   *slog.Logger
}

func (s *Server) Register(fn func(*grpc.Server)) *Server {
	fn(s.core)
	return s
}

func (s *Server) Start(ctx context.Context) error {
	s.logger.LogAttrs(ctx, slog.LevelInfo, "starting",
		slog.String("address", s.listener.Addr().String()),
	)
	if err := s.core.Serve(s.listener); err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) {
	s.logger.LogAttrs(ctx, slog.LevelInfo, "stopped")
	s.core.GracefulStop()
}

func (s *Server) Listen(ctx context.Context) {
	go func() {
		if err := s.Start(ctx); err != nil {
			s.logger.LogAttrs(ctx, slog.LevelError, err.Error())
		}
	}()

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM)
	stopSignal := <-stopCh

	s.logger.LogAttrs(ctx, slog.LevelInfo, "received stop signal",
		slog.Any("signal", stopSignal),
	)
	s.Stop(ctx)
}
