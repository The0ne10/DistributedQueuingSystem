package main

import (
	"DistributedQueueSystem/internal/config"
	"log/slog"
	"os"
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

	// TODO: init storage

	// TODO: init grpcServer

	// TODO: init application
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
