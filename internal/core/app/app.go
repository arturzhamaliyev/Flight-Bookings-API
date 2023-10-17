package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/core/logging"
	"go.uber.org/zap"
)

// Listener represents a type that can listen for incoming connections.
type Listener interface {
	Listen(ctx context.Context) error
}

// OnShutdownFunc is a function that is called when the app is shutdown.
type OnShutdownFunc func()

// App represents application run by this service.
type App struct {
	Name          string
	shutdownFuncs []OnShutdownFunc
}

// Launch is a function that is called when the app is started.
type Launch func(ctx context.Context, a *App) ([]Listener, error)

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

	listeners, err := launch(ctx, a)
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

	var wg sync.WaitGroup

	for _, listener := range listeners {
		wg.Add(1)
		listener := listener

		go func() {
			defer wg.Done()

			err := listener.Listen(ctx)
			if err != nil {
				logging.From(ctx).Error("listener failed", zap.Error(err))
			}
		}()
	}
	wg.Wait()

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
