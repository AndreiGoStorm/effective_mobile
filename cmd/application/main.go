package main

import (
	"context"
	_ "effective_mobile/cmd/application/docs"
	"effective_mobile/internal/config"
	"effective_mobile/internal/logger"
	"effective_mobile/internal/middlewares"
	"effective_mobile/internal/route"
	"effective_mobile/internal/storage"
	"encoding/json"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

func main() {
	cfg := config.New()

	log := logger.New(&cfg.Logger)

	log.Info("starting application",
		slog.String("name", cfg.Name), slog.String("version", cfg.Version))

	ctx := context.Background()
	store := storage.NewStorage(cfg.Database)
	if err := store.Connect(ctx); err != nil {
		log.Error("failed store connect", logger.Err(err))
		os.Exit(1)
	}

	app := BuildServer(store, log)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt)

	go func() {
		if err := app.Listen(net.JoinHostPort(cfg.HTTPServer.Host, strconv.Itoa(cfg.HTTPServer.Port))); err != nil {
			log.Error("failed to start server", logger.Err(err))
		}
	}()

	<-stop
	err := store.Close(ctx)
	if err != nil {
		log.Error("failed store close", logger.Err(err))
	}

	log.Info("Gracefully stopped")
}

func BuildServer(store *storage.Storage, log *slog.Logger) (app *fiber.App) {
	app = fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(recover.New())
	app.Use(middlewares.Cors())

	apiHandlers := route.New(store, log)
	route.V1Routes(app, apiHandlers)
	app.Get("/swagger/*", swagger.HandlerDefault)

	return
}
