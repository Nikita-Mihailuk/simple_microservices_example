package user

import (
	"context"
	"github.com/Nikita-Mihailuk/simple_microservices_example/user_service/internal/domain/model"
	userServicev1 "github.com/Nikita-Mihailuk/simple_microservices_example/user_service/protoc/gen/go/user_service"
	"google.golang.org/grpc"
)

type UserService interface {
	CreateNewUser(ctx context.Context, name string, age int32) (int64, error)
	GetUsers(ctx context.Context) ([]model.User, error)
}

type serverGRPC struct {
	userService UserService
	userServicev1.UnimplementedUserServer
}

func RegisterGRPCServer(grpcServer *grpc.Server, userService UserService) {
	userServicev1.RegisterUserServer(grpcServer, &serverGRPC{userService: userService})
}

func (s *serverGRPC) CreateUser(ctx context.Context, req *userServicev1.CreateUserRequest) (*userServicev1.CreateUserResponse, error) {
	userID, err := s.userService.CreateNewUser(ctx, req.GetName(), req.GetAge())
	if err != nil {
		return nil, err
	}
	return &userServicev1.CreateUserResponse{UserId: userID}, nil
}

func (s *serverGRPC) GetAllUsers(ctx context.Context, req *userServicev1.GetAllUsersRequest) (*userServicev1.GetAllUsersResponse, error) {
	users, err := s.userService.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	pbUsers := make([]*userServicev1.UserData, 0, len(users))
	for _, user := range users {
		pbUsers = append(pbUsers, &userServicev1.UserData{
			UserId: user.ID,
			Name:   user.Name,
			Age:    user.Age,
		})
	}

	return &userServicev1.GetAllUsersResponse{Users: pbUsers}, nil
}
