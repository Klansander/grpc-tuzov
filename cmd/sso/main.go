package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
	"sso/internal/config"
	"strings"
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

	logrus.Info("Инициализация конфигурации")
	ctx = config.ContextWithConfig(ctx, config.MustLoad())

	logrus.Info("Инициализация App")
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
