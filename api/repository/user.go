package repository

import (
	"blog/infrastructure"
	"blog/models"
	"blog/util"
)

//UserRepository -> UserRepository resposible for accessing database
type UserRepository struct {
	db infrastructure.Database
}

//NewUserRepository -> creates a instance on UserRepository
func NewUserRepository(db infrastructure.Database) UserRepository {
	return UserRepository{
		db: db,
	}
}

//CreateUser -> method for saving user to database
func (u UserRepository) CreateUser(user models.UserRegister) error {

	var dbUser models.User
	dbUser.Email = user.Email
	dbUser.FirstName = user.FirstName
	dbUser.LastName = user.LastName
	dbUser.Password = user.Password
	dbUser.IsActive = false
	return u.db.DB.Create(&dbUser).Error
}

//LoginUser -> method for returning user
func (u UserRepository) LoginUser(user models.UserLogin) (*models.User, error) {

	var dbUser models.User
	email := user.Email
	password := user.Password

	err := u.db.DB.Where("email = ?", email).First(&dbUser).Error
	if err != nil {
		return nil, err
	}

	hashErr := util.CheckPasswordHash(password, dbUser.Password)
	if hashErr != nil {
		return nil, hashErr
	}
	return &dbUser, nil
}
