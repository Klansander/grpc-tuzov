package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"log/slog"
	"os"
	"os/signal"
	"path"
	"runtime"
	"sso/internal/app"
	"sso/internal/config"
	"sso/model"
	"strings"
	"syscall"
)

func init() {

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
		DisableColors:   false,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			_, fileName := path.Split(f.File)
			arr := strings.Split(f.Function, ".")

			dir := " " + arr[0] + "/"
			funcName := strings.Join(arr[1:], ".")

			return funcName, dir + fileName
		},
	})

	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(true)

}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cfg := config.MustLoad()
	logrus.Info("Инициализация конфигурации")
	ctx = config.ContextWithConfig(ctx, config.MustLoad())

	logrus.Info("Инициализация App")

	//log := setupLogger(cfg.Env)
	//
	//log.Info("Старт Приложения",
	//	slog.String("env", cfg.Env),
	//	slog.Any("cfg", cfg),
	//	slog.Int("port", cfg.GRPC.Port),
	//)

	applications := app.New(cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)
	go applications.GRPCSrv.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	logrus.Info(<-stop)

	logrus.Info("Остановка к приложению")

	applications.GRPCSrv.Stop()

	logrus.Info("Приложение остановлено")

	//a, err := app.NewApp(ctx)
	//if err != nil {
	//	logrus.Fatal(err)
	//}
	//
	//g, ctx := errgroup.WithContext(ctx)
	//
	//g.Go(func() error {
	//
	//	signalChannel := make(chan os.Signal, 1)
	//	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	//
	//	select {
	//	case <-signalChannel:
	//		cancel()
	//
	//	case <-ctx.Done():
	//		return ctx.Err()
	//	}
	//
	//	a.Stop(ctx)
	//
	//	return nil
	//
	//})
	//
	//g.Go(func() error {
	//
	//	logrus.Info("Запуск App")
	//
	//	return a.Run(ctx)
	//
	//})
	//
	//if err := g.Wait(); err != nil {
	//	logrus.Errorf("Приложение упало с ошибкой: %v", err)
	//}
	//
	//logrus.Warn("app stopped")

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {

	case model.EnvLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case model.EnvDev:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case model.EnvProd:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	}
	return log
}
