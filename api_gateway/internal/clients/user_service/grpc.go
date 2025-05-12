package user_service

import (
	"context"
	"github.com/Nikita-Mihailuk/simple_microservices_example/api_gateway/internal/domain/model"
	userServicev1 "github.com/Nikita-Mihailuk/simple_microservices_example/api_gateway/protoc/gen/go/user_service"
	"google.golang.org/grpc"
)

type UserClient struct {
	api userServicev1.UserClient
}

func NewUserClient(ctx context.Context, addr string) (*UserClient, error) {
	cc, err := grpc.DialContext(ctx, addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &UserClient{api: userServicev1.NewUserClient(cc)}, nil
}

func (u *UserClient) CreateUser(ctx context.Context, name string, age int32) (int64, error) {
	resp, err := u.api.CreateUser(ctx, &userServicev1.CreateUserRequest{
		Name: name,
		Age:  age,
	})
	if err != nil {
		return 0, err
	}

	return resp.GetUserId(), nil
}

func (u *UserClient) GetUsers(ctx context.Context) ([]model.User, error) {
	resp, err := u.api.GetAllUsers(ctx, &userServicev1.GetAllUsersRequest{})
	if err != nil {
		return nil, err
	}

	users := make([]model.User, 0, len(resp.Users))
	for _, pbUser := range resp.Users {
		users = append(users, model.User{
			ID:   pbUser.UserId,
			Name: pbUser.Name,
			Age:  pbUser.Age,
		})
	}

	return users, nil
}
