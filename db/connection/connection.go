package connection

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDb(host string, user string, password string, database string, port string) (*gorm.DB, error)  {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
			host, user, password, database, port,
		)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
