package user

import (
	"context"

	userModel "github.com/anditakaesar/uwa-back/adapter/models/user"
)

type UserSeviceInterface interface {
	CreateUser(ctx context.Context, user *userModel.User) error
	GetUserByUsername(ctx context.Context, username string) (*userModel.User, error)
}

type UserService struct {
	UserModel userModel.UserModelInterface
}

func NewUserService(userModel userModel.UserModelInterface) UserSeviceInterface {
	return &UserService{
		UserModel: userModel,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *userModel.User) error {
	return s.UserModel.CreateUser(ctx, user)
}

func (s *UserService) GetUserByUsername(ctx context.Context, username string) (*userModel.User, error) {
	return s.UserModel.GetUserByUsername(ctx, username)
}
