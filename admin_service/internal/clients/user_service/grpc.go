package user_service

import (
	"context"
	"github.com/Nikita-Mihailuk/simple_microservices_example/admin_service/internal/domain/model"
	userServicev1 "github.com/Nikita-Mihailuk/simple_microservices_example/admin_service/protoc/gen/go/user_service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Client struct {
	api userServicev1.UserClient
	log *zap.Logger
}

func NewClient(ctx context.Context, log *zap.Logger, addr string) (*Client, error) {
	cc, err := grpc.DialContext(ctx, addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Client{api: userServicev1.NewUserClient(cc), log: log}, nil
}

func (c *Client) GetUsers(ctx context.Context) ([]model.User, error) {
	resp, err := c.api.GetAllUsers(ctx, &userServicev1.GetAllUsersRequest{})
	if err != nil {
		return nil, err
	}

	pbUsers := resp.Users
	users := make([]model.User, 0, len(pbUsers))
	for _, pbUser := range pbUsers {
		users = append(users, model.User{
			ID:   pbUser.UserId,
			Name: pbUser.Name,
			Age:  pbUser.Age,
		})
	}

	return users, nil
}
