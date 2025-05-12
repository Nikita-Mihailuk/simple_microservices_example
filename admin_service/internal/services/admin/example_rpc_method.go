package admin

import (
	"context"
	"github.com/Nikita-Mihailuk/simple_microservices_example/admin_service/internal/clients/user_service"
	"github.com/Nikita-Mihailuk/simple_microservices_example/admin_service/internal/domain/model"
	"go.uber.org/zap"
)

// example method for RPC
func GetUsers(log *zap.Logger, userServiceClient *user_service.Client) ([]model.User, error) {
	users, err := userServiceClient.GetUsers(context.TODO())
	if err != nil {
		return nil, err
	}

	log.Info("get users from user service", zap.Any("users", users))
	return users, nil
}
