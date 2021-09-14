package app

import (
	"github.com/Kortex/golang-bookstore-users/db/connection"
	"github.com/Kortex/golang-bookstore-users/db/entities"
	"github.com/Kortex/golang-bookstore-users/utils/constants"
	"github.com/gin-gonic/gin"
	"log"
)

var router = gin.Default()

func StartApplication()  {
	host, user, password, database, port := "localhost", "user", "password", "database", "5432"
	db, err := connection.ConnectToDb(host, user, password, database, port)
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&entities.UserEntity{})
	InitRoutes()
	err = router.Run(constants.ServerPort)
	if nil != err {
		log.Fatalln(err)
	}
	log.Printf("Successfully started web server on port %s", constants.ServerPort)
}
