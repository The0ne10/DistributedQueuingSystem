package app

import (
	"DistributedQueueSystem/internal/app/grpcapp"
	"DistributedQueueSystem/internal/config"
	"DistributedQueueSystem/internal/services"
	"fmt"
	"log/slog"
)

type App struct {
	GRPCSrv  *grpcapp.App
	Services *services.ServiceContainer
}

func New(
	cfg config.Config,
	log *slog.Logger,
	gRPCPort int,
) (*App, error) {
	// TODO: пробросить storage в сервис контейнер

	servicesContainer, err := services.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize services: %w", err)
	}

	grpcApp := grpcapp.New(log, gRPCPort, servicesContainer)

	return &App{
		GRPCSrv: grpcApp,
	}, nil
}
