package users

import (
	"github.com/Kortex/golang-bookstore-users/models/users"
	"github.com/Kortex/golang-bookstore-users/services"
	"github.com/Kortex/golang-bookstore-users/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.BadRequest("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusOK, &result)
}

func GetUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {

	}

	c.JSON(http.StatusNotImplemented, "implement me!")
}

func SearchUser(c *gin.Context)  {
	c.String(http.StatusNotImplemented, "implement me!")
}
