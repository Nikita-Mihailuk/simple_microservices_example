package admin_service

import (
	"context"
	adminServicev1 "github.com/Nikita-Mihailuk/simple_microservices_example/api_gateway/protoc/gen/go/admin_service"
	"google.golang.org/grpc"
)

type AdminClient struct {
	api adminServicev1.AdminClient
}

func NewAdminClient(ctx context.Context, addr string) (*AdminClient, error) {
	cc, err := grpc.DialContext(ctx, addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &AdminClient{api: adminServicev1.NewAdminClient(cc)}, nil
}

func (a *AdminClient) DeleteUserById(ctx context.Context, userID int64) (bool, error) {
	resp, err := a.api.DeleteUser(ctx, &adminServicev1.DeleteUserRequest{UserId: userID})
	if err != nil {
		return false, err
	}
	return resp.GetSuccess(), nil
}
