package app

import (
	"DistributedQueueSystem/internal/app/grpcapp"
	"DistributedQueueSystem/internal/services"
	"log/slog"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	gRPCPort int,
) *App {
	servicesContainer := services.New()

	grpcApp := grpcapp.New(log, servicesContainer, gRPCPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
