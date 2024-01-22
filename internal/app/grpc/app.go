package grpcapp

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log/slog"
	"net"
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

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	logrus.Infof("op", "grpcapp.Run")
	logrus.Infof("port", a.port)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		fmt.Println("Error run app")
	}

	logrus.Infof("servis runn", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		fmt.Errorf("%s:%w", "grpcapp", err)
	}
	return nil
}

func (a *App) Stop() {
	a.gRPCServer.GracefulStop()
}
