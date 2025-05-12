package grpc

import (
	"fmt"
	userGRPC "github.com/Nikita-Mihailuk/simple_microservices_example/user_service/internal/delivery/grpc/user"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type App struct {
	log        *zap.Logger
	grpcServer *grpc.Server
	port       int
}

func NewApp(log *zap.Logger, userService userGRPC.UserService, port int) *App {
	grpcServer := grpc.NewServer()
	userGRPC.RegisterGRPCServer(grpcServer, userService)
	return &App{log, grpcServer, port}
}

func (a *App) Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		panic(err)
	}

	a.log.Info("starting grpc server", zap.String("address", lis.Addr().String()))

	if err = a.grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
func (a *App) Stop() {
	a.log.Info("stopping grpc server")
	a.grpcServer.GracefulStop()
}
