package admin

import (
	"context"
	"github.com/Nikita-Mihailuk/simple_microservices_example/admin_service/internal/clients/user_service"
	"go.uber.org/zap"
)

type AdminService struct {
	log               *zap.Logger
	adminDeleter      AdminDeleter
	userServiceClient *user_service.Client
}

func NewAdminService(log *zap.Logger, adminDeleter AdminDeleter, userServiceClient *user_service.Client) *AdminService {
	return &AdminService{
		log:               log,
		adminDeleter:      adminDeleter,
		userServiceClient: userServiceClient,
	}
}

type AdminDeleter interface {
	DeleteUserByID(ctx context.Context, userID int64) (bool, error)
}

func (a *AdminService) DeleteUser(ctx context.Context, userID int64) (bool, error) {
	isDeleted, err := a.adminDeleter.DeleteUserByID(ctx, userID)
	if err != nil {
		return false, err
	}

	a.log.Info("deleted user", zap.Int64("user_id", userID))
	return isDeleted, err
}
