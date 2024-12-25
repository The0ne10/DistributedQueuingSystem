package main

import (
	"DistributedQueueSystem/internal/app"
	"DistributedQueueSystem/internal/config"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (
	localEnv = "local"
	devEnv   = "dev"
	prodEnv  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("Logger starting")

	if cfg.Env == localEnv {
		log.Debug("Debug starting")
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	application := app.New(log, cfg.GRPC.Port)

	go application.GRPCSrv.MustRun()

	<-sigChan

	log.Info("Application stopped.")

	// TODO: Сделать мониторинг сервиса
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case localEnv:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case devEnv:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case prodEnv:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
