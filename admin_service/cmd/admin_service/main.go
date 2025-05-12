package main

import (
	"fmt"
	"github.com/Nikita-Mihailuk/simple_microservices_example/admin_service/internal/app"
	"github.com/Nikita-Mihailuk/simple_microservices_example/admin_service/internal/config"
	"github.com/Nikita-Mihailuk/simple_microservices_example/admin_service/pkg/logging"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.GetConfig()
	log := logging.GetLogger()

	userServiceAddr := fmt.Sprintf("%s:%s", cfg.UserService.Host, cfg.UserService.Port)

	application := app.NewApp(log, cfg.GRPC.Port, cfg.UserService.Timeout, userServiceAddr)

	go application.GRPCServer.Run()

	// Graceful shutdown

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	application.GRPCServer.Stop()
	log.Info("application stopped")
}
