package main

import (
	"fmt"
	"github.com/Nikita-Mihailuk/simple_microservices_example/api_gateway/internal/app"
	"github.com/Nikita-Mihailuk/simple_microservices_example/api_gateway/internal/config"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.GetConfig()

	userServiceAddr := fmt.Sprintf("%s:%s", cfg.UserService.Host, cfg.UserService.Port)
	adminServiceAddr := fmt.Sprintf("%s:%s", cfg.AdminService.Host, cfg.AdminService.Port)

	application := app.NewApp(
		cfg.HTTPServer.Port,
		cfg.UserService.Timeout,
		userServiceAddr,
		cfg.AdminService.Timeout,
		adminServiceAddr,
	)

	go application.HTTPServer.Run()

	// Graceful shutdown

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	application.HTTPServer.Stop()
}
