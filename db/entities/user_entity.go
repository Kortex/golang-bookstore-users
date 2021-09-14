package entities

import "gorm.io/gorm"

type UserEntity struct {
	gorm.Model
	FirstName string
	LastName string
	Email string
}

func (UserEntity) TableName() string {
	return "gorm_users"
}

