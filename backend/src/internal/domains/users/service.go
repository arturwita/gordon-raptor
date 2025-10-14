package users

import "context"

type UserService interface {
	CreateUser(dto *CreateUserDto, ctx context.Context) (*UserModel, error)
	GetUserByEmail(email string, ctx context.Context) (*UserModel, error)
}

type userService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) (UserService, error) {
	return &userService{repository}, nil
}

func (service *userService) CreateUser(dto *CreateUserDto, ctx context.Context) (*UserModel, error) {
	return service.repository.CreateUser(dto, ctx)
}

func (service *userService) GetUserByEmail(email string, ctx context.Context) (*UserModel, error) {
	return service.repository.GetUserByEmail(email, ctx)
}
