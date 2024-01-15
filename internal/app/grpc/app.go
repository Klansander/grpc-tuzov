package grpcapp

import (
	"google.golang.org/grpc"
	authgrpc "sso/internal/grpc/auth"
)

type App struct {
	gRPCServer *grpc.Server
	port       int
}

func New(port int) *App {
	gPRCServer := grpc.NewServer()
	authgrpc.NewServer(gPRCServer)
	return &App{
		gRPCServer: gPRCServer,
		port:       port,
	}

}

func (a *App) Run() error {
	
}
