package services

import (
	"github.com/Kortex/golang-bookstore-users/db/entities"
	"github.com/Kortex/golang-bookstore-users/models/users"
	"github.com/Kortex/golang-bookstore-users/utils/errors"
	"gorm.io/gorm"
)

var DB *gorm.DB

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	DB.Save(&entities.UserEntity{
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
	})
	return &user, nil
}
