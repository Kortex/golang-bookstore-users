package app

import (
	"github.com/Kortex/golang-bookstore-users/controllers/ping"
	"github.com/Kortex/golang-bookstore-users/controllers/users"
)

func InitRoutes() {

	// health endpoint
	router.GET("/ping", ping.Ping)

	// users mappings
	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/search", users.SearchUser)

}
