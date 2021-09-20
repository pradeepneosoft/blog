package service

import (
	"blog/api/repository"
	"blog/models"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}
func (u UserService) CreateUser(user models.UserRegister) error {
	return u.repo.CreateUser(user)
}

func (u UserService) LoginUser(user models.UserLogin) (*models.User, error) {
	return u.repo.LoginUser(user)
}
func (u UserService) FindAllUser(user models.User, keyword string) (*[]models.User, int64, error) {
	return u.repo.FindAllUser(user, keyword)
}
