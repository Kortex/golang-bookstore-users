package props

import (
	"log"
	"os"
	"reflect"
	"strings"
)

const webPort string = "WEB_PORT"
const hostAddress string = "HOST_ADDRESS"

var WebProps ServerProps

type ServerProps struct {
	address string
	port    string
}

func (s ServerProps) Address() string {
	return s.address
}

func (s ServerProps) Port() string {
	return s.port
}

func (s ServerProps) validate() {
	fields := reflect.TypeOf(s)
	values := reflect.ValueOf(s)
	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		value := values.Field(i)
		if trimmed := strings.TrimSpace(value.String()); trimmed == "" {
			log.Fatalf("Please provide a valid value for server props field %s", field.Name)
		}
	}
}

func init() {
	WebProps = ServerProps{
		port:    os.Getenv(webPort),
		address: os.Getenv(hostAddress),
	}
	WebProps.validate()
}
