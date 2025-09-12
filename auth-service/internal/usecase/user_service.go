package usecase

import (
	"github.com/MateusQ7/ecommerce-go/auth-service/internal/domain"
	"github.com/MateusQ7/ecommerce-go/auth-service/pkg/utils"
)

type UserRepository interface {
	CreateUser(user *domain.User) error
	FindAll() ([]domain.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) *UserService {
	return &UserService{repo: r}
}

func (us *UserService) CreateNewUser(user *domain.User) error {
	hashedPassowrd, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	user.Password = hashedPassowrd
	return us.repo.CreateUser(user)
}

func (us *UserService) ListUsers() ([]domain.User, error) {
	return us.repo.FindAll()
}
