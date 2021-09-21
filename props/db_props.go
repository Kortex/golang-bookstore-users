package props

import (
	"log"
	"os"
	"reflect"
	"strings"
)

const databaseHost = "DATABASE_HOST"
const databaseUser = "DATABASE_USER"
const databasePassword = "DATABASE_PASSWORD"
const databaseName = "DATABASE_NAME"
const databasePort = "DATABASE_PORT"

type DbProps struct {
	host     string
	user     string
	password string
	database string
	port     string
}

func (c DbProps) Host() string {
	return c.host
}

func (c DbProps) User() string {
	return c.user
}

func (c DbProps) Password() string {
	return c.password
}

func (c DbProps) Database() string {
	return c.database
}

func (c DbProps) Port() string {
	return c.port
}

func (c DbProps) validate() {
	fields := reflect.TypeOf(c)
	values := reflect.ValueOf(c)
	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		value := values.Field(i)
		if trimmed := strings.TrimSpace(value.String()); trimmed == "" {
			log.Fatalf("Please provide a valid value for database props field %s", field.Name)
		}
	}
}

var DBProps DbProps

func init() {
	DBProps = DbProps{
		host:     os.Getenv(databaseHost),
		user:     os.Getenv(databaseUser),
		password: os.Getenv(databasePassword),
		database: os.Getenv(databaseName),
		port:     os.Getenv(databasePort),
	}
	DBProps.validate()
}
