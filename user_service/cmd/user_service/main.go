package main

import (
	"github.com/Nikita-Mihailuk/simple_microservices_example/user_service/internal/app"
	"github.com/Nikita-Mihailuk/simple_microservices_example/user_service/internal/config"
	"github.com/Nikita-Mihailuk/simple_microservices_example/user_service/pkg/logging"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.GetConfig()
	log := logging.GetLogger()

	application := app.NewApp(log, cfg.GRPC.Port)

	go application.GRPCServer.Run()

	// Graceful shutdown

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	application.GRPCServer.Stop()
	log.Info("application stopped")
}
