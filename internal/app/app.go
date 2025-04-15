package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/MaxFando/structura/internal/server"
	v1 "github.com/MaxFando/structura/internal/server/structura/v1"
	"github.com/MaxFando/structura/pkg/log"
)

type App struct {
	srv    *server.Server
	logger log.Logger
}

func New() *App {
	logger := log.NewLogger()

	return &App{
		logger: logger.With("app", "structura"),
	}
}

func (a *App) Logger() log.Logger {
	return a.logger
}

func (a *App) Init(ctx context.Context) error {
	return nil
}

func (a *App) Run(ctx context.Context) error {
	structuraServer := v1.NewStructuraServer()

	srv, err := server.NewServer(a.logger, structuraServer)
	if err != nil {
		return fmt.Errorf("ошибка при инициализации сервера: %w", err)
	}
	srv.Serve(ctx)

	a.srv = srv

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGTERM, syscall.SIGINT)

	select {
	case s := <-interrupt:
		return fmt.Errorf("signal: %s", s.String())
	case err := <-srv.Notify():
		return fmt.Errorf("server: %w", err)
	}
}

func (a *App) Shutdown(ctx context.Context) {
	a.srv.Shutdown(ctx)
}
