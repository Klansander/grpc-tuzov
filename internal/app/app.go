package app

import (
	grpcapp "sso/internal/app/grpc"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	grpcApp := grpcapp.New(grpcPort)
	return &App{GRPCSrv: grpcApp}
}
