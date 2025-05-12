package user

import (
	"context"
	"github.com/Nikita-Mihailuk/simple_microservices_example/user_service/internal/domain/model"
	"go.uber.org/zap"
)

type UserService struct {
	log          *zap.Logger
	userSaver    UserSaver
	userProvider UserProvider
}

func NewUserService(log *zap.Logger, userSaver UserSaver, userProvider UserProvider) *UserService {
	return &UserService{
		log:          log,
		userSaver:    userSaver,
		userProvider: userProvider,
	}
}

type UserSaver interface {
	SaveUser(ctx context.Context, name string, age int32) (int64, error)
}

type UserProvider interface {
	GetUsers(ctx context.Context) ([]model.User, error)
}

func (u *UserService) CreateNewUser(ctx context.Context, name string, age int32) (int64, error) {
	userID, err := u.userSaver.SaveUser(ctx, name, age)
	if err != nil {
		return 0, err
	}
	u.log.Info("user created", zap.Int64("user_id", userID))
	return userID, nil
}
func (u *UserService) GetUsers(ctx context.Context) ([]model.User, error) {
	users, err := u.userProvider.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	u.log.Info("users retrieved")
	return users, nil
}
