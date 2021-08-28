package service

import (
	"blog/api/repository"
	"blog/models"
)

//UserService UserService struct
type UserService struct {
	repo repository.UserRepository
}

//NewUserService : get injected user repo
func NewUserService(repo repository.UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

//Save -> saves users entity
func (u UserService) CreateUser(user models.UserRegister) error {
	return u.repo.CreateUser(user)
}

//Login -> Gets validated user
func (u UserService) LoginUser(user models.UserLogin) (*models.User, error) {
	return u.repo.LoginUser(user)

}
