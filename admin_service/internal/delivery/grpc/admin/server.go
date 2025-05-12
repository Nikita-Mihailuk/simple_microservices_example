package admin

import (
	"context"
	adminServicev1 "github.com/Nikita-Mihailuk/simple_microservices_example/admin_service/protoc/gen/go/admin_service"
	"google.golang.org/grpc"
)

type AdminService interface {
	DeleteUser(ctx context.Context, userID int64) (bool, error)
}

type serverGRPC struct {
	adminService AdminService
	adminServicev1.UnimplementedAdminServer
}

func RegisterGRPCServer(grpcServer *grpc.Server, userService AdminService) {
	adminServicev1.RegisterAdminServer(grpcServer, &serverGRPC{adminService: userService})
}

func (s *serverGRPC) DeleteUser(ctx context.Context, req *adminServicev1.DeleteUserRequest) (*adminServicev1.DeleteUserResponse, error) {
	isDeleted, err := s.adminService.DeleteUser(ctx, req.GetUserId())
	if err != nil {
		return &adminServicev1.DeleteUserResponse{Success: isDeleted}, err
	}

	return &adminServicev1.DeleteUserResponse{Success: isDeleted}, nil
}
