package app

import (
	"DistributedQueueSystem/internal/app/grpcapp"
	"DistributedQueueSystem/internal/services"
	"log/slog"
)

type App struct {
	GRPCSrv  *grpcapp.App
	Services *services.ServiceContainer
}

func New(
	log *slog.Logger,
	gRPCPort int,
) *App {
	// TODO: пробросить storage в сервис контейнер

	servicesContainer := services.New()

	grpcApp := grpcapp.New(log, gRPCPort, servicesContainer)

	return &App{
		GRPCSrv: grpcApp,
	}
}
