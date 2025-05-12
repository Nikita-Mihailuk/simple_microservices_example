package app

import (
	"context"
	grpcApp "github.com/Nikita-Mihailuk/simple_microservices_example/user_service/internal/app/grpc"
	userService "github.com/Nikita-Mihailuk/simple_microservices_example/user_service/internal/services/user"
	"github.com/Nikita-Mihailuk/simple_microservices_example/user_service/internal/storage/postgres"
	"go.uber.org/zap"
)

type App struct {
	GRPCServer *grpcApp.App
}

func NewApp(log *zap.Logger, port int) *App {
	postgresClient := postgres.NewClient(context.TODO())
	storage := postgres.NewStorage(postgresClient)

	service := userService.NewUserService(log, storage, storage)

	gRPCApp := grpcApp.NewApp(log, service, port)
	return &App{GRPCServer: gRPCApp}
}
