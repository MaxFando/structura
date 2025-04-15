package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/MaxFando/structura/internal/app"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	application := app.New()
	err := application.Init(ctx)
	if err != nil {
		application.Logger().Error(ctx, "Application initialization failed", "error", err)
		return
	}

	if err := application.Run(ctx); err != nil {
		application.Logger().Error(ctx, "Application run failed", "error", err)
	}

	cancel()
	application.Logger().Info(ctx, "Application finished")
}
