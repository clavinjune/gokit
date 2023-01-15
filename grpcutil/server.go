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
	logger   *slog.Logger
	core     *grpc.Server
	Listener net.Listener
}

func (s *Server) Register(fn func(*grpc.Server)) *Server {
	fn(s.core)
	return s
}

func (s *Server) Start(_ context.Context) error {
	s.logger.LogAttrs(slog.LevelInfo, "starting",
		slog.String("address", s.Listener.Addr().String()),
	)
	if err := s.core.Serve(s.Listener); err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop(_ context.Context) {
	s.logger.LogAttrs(slog.LevelInfo, "stopped")
	s.core.GracefulStop()
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
