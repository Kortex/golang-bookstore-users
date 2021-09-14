package constants

import (
	"os"
)

var (
	ServerPort = os.Getenv("USER_API_SERVER_PORT")
)