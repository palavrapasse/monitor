package http

import (
	"fmt"
	"os"
)

const SantosServiceName = "santos"

const (
	santosServerHostEnvKey = "santos_server_host"
	santosServerPortEnvKey = "santos_server_port"
)

var (
	santosServerHost            = os.Getenv(santosServerHostEnvKey)
	santosServerPort            = os.Getenv(santosServerPortEnvKey)
	santosServiceHealthEndpoint = "health"
	SantosServiveHealthURL      = fmt.Sprintf(URL_FORMAT, santosServerHost, santosServerPort, santosServiceHealthEndpoint)
)
