package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/core/logging"

	"go.uber.org/zap"
)

// OnShutdownFunc is a function that is called when the app is shutdown.
type OnShutdownFunc func()

// App represents application run by this service.
type App struct {
	Name          string
	shutdownFuncs []OnShutdownFunc
}

// Launch is a function that is called when the app is started.
type Launch func(ctx context.Context, a *App) (func(), error)

// Start starts the application.
func Start(launch Launch) {
	ctx := context.Background()

	a := &App{
		Name: "flight booking api",
	}

	a.OnShutdown(func() {
		if err := logging.Sync(ctx); err != nil {
			log.Printf("failed to sync logging: %v", err)
		}
	})

	logging.From(ctx).Info("app starting...")

	listenAndServe, err := launch(ctx, a)
	if err != nil {
		logging.From(ctx).Fatal("failed to start app", zap.Error(err))
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		shutdown(ctx, a)
		os.Exit(1)
	}()

	listenAndServe()

	shutdown(ctx, a)
}

// OnShutdown registers a function that is called when the app is shutdown.
func (a *App) OnShutdown(f func()) {
	a.shutdownFuncs = append([]OnShutdownFunc{f}, a.shutdownFuncs...)
}

func shutdown(ctx context.Context, a *App) {
	for _, shutdownFunc := range a.shutdownFuncs {
		shutdownFunc()
	}

	logging.From(ctx).Info("app shutdown")
}
