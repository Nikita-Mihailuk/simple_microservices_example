package app

import (
	"context"
	grpcApp "github.com/Nikita-Mihailuk/simple_microservices_example/admin_service/internal/app/grpc"
	"github.com/Nikita-Mihailuk/simple_microservices_example/admin_service/internal/clients/user_service"
	adminService "github.com/Nikita-Mihailuk/simple_microservices_example/admin_service/internal/services/admin"
	"github.com/Nikita-Mihailuk/simple_microservices_example/admin_service/internal/storage/postgres"
	"go.uber.org/zap"
	"time"
)

type App struct {
	GRPCServer *grpcApp.App
}

func NewApp(log *zap.Logger, port int, timeout time.Duration, userServiceAddr string) *App {
	postgresClient := postgres.NewClient(context.TODO())
	storage := postgres.NewStorage(postgresClient)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	userServiceClient, err := user_service.NewClient(ctx, log, userServiceAddr)
	if err != nil {
		panic(err)
	}

	// example RPC
	// adminService.GetUsers(log, userServiceClient)

	service := adminService.NewAdminService(log, storage, userServiceClient)

	gRPCApp := grpcApp.NewApp(log, service, port)
	return &App{GRPCServer: gRPCApp}
}
