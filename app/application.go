package app

import (
	"fmt"
	"github.com/Kortex/golang-bookstore-users/db/connection"
	"github.com/Kortex/golang-bookstore-users/db/entities"
	"github.com/Kortex/golang-bookstore-users/props"
	"github.com/gin-gonic/gin"
	"log"
)

var router = gin.Default()

func StartApplication() {
	db, err := connection.ConnectToDb(
		props.DBProps.Host(),
		props.DBProps.User(),
		props.DBProps.Password(),
		props.DBProps.Database(),
		props.DBProps.Port())
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&entities.UserEntity{})
	InitRoutes()
	err = router.Run(fmt.Sprintf("%s%s", props.WebProps.Address(), props.WebProps.Port()))
	if nil != err {
		log.Fatalln(err)
	}
	log.Printf("Successfully started web server on port %s", props.WebProps.Port())
}
