package grpcutil

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

type Proxy struct {
	_          struct{}
	logger     *slog.Logger
	core       *http.Server
	proxy      *runtime.ServeMux
	Listener   net.Listener
	grpcServer *Server
}

func (p *Proxy) Register(fn func(server *grpc.Server, proxy *runtime.ServeMux, endpoint string)) *Proxy {
	fn(p.grpcServer.core, p.proxy, p.grpcServer.Listener.Addr().String())
	return p
}

func (p *Proxy) Start(ctx context.Context) error {
	go func() {
		if err := p.grpcServer.Start(ctx); err != nil {
			p.logger.LogAttrs(slog.LevelError, err.Error())
		}
	}()

	p.logger.LogAttrs(slog.LevelInfo, "starting",
		slog.String("address", p.Listener.Addr().String()),
	)
	if err := p.core.Serve(p.Listener); err != nil {
		return err
	}

	return nil
}

func (p *Proxy) Stop(ctx context.Context) {
	p.logger.LogAttrs(slog.LevelInfo, "stopped")
	p.core.Shutdown(ctx)
	p.grpcServer.Stop(ctx)
}

func (p *Proxy) Listen(ctx context.Context) {
	go func() {
		if err := p.Start(ctx); !errors.Is(err, http.ErrServerClosed) {
			p.logger.LogAttrs(slog.LevelError, err.Error())
		}
	}()

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM)
	stopSignal := <-stopCh

	p.logger.LogAttrs(slog.LevelInfo, "received stop signal",
		slog.Any("signal", stopSignal),
	)
	p.Stop(ctx)
}
