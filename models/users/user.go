package users

import (
	"github.com/Kortex/golang-bookstore-users/utils/errors"
	"strings"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
}
func (user *User) Validate() *errors.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.BadRequest("invalid email address")
	}
	return nil
}