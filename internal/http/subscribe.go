package http

import (
	"fmt"
	"os"
)

const SubscribeServiceName = "subscribe"

const (
	subscribeServerHostEnvKey = "subscribe_server_host"
	subscribeServerPortEnvKey = "subscribe_server_port"
)

var (
	subscribeServerHost            = os.Getenv(subscribeServerHostEnvKey)
	subscribeServerPort            = os.Getenv(subscribeServerPortEnvKey)
	subscribeServiceHealthEndpoint = "health"
	SubscribeServiveHealthURL      = fmt.Sprintf(URL_FORMAT, subscribeServerHost, subscribeServerPort, subscribeServiceHealthEndpoint)
)
