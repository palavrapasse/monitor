package http

import (
	"fmt"
	"os"
)

const (
	queryServerHostEnvKey = "query_server_host"
	queryServerPortEnvKey = "query_server_port"
)

var (
	queryServerHost            = os.Getenv(queryServerHostEnvKey)
	queryServerPort            = os.Getenv(queryServerPortEnvKey)
	queryServiceHealthEndpoint = "/health"
	QueryServiveHealthURL      = fmt.Sprintf(URL_FORMAT, queryServerHost, queryServerPort, queryServiceHealthEndpoint)
)
