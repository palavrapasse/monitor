package http

import (
	"fmt"
	"os"
)

const (
	subscribeServerHostEnvKey = "subscribe_server_host"
	subscribeServerPortEnvKey = "subscribe_server_port"
)

var (
	subscribeServerHost            = os.Getenv(subscribeServerHostEnvKey)
	subscribeServerPort            = os.Getenv(subscribeServerPortEnvKey)
	subscribeServiceHealthEndpoint = "health"
	SubscribeServiveHealthURL      = fmt.Sprintf("%s:%s/%s", subscribeServerHost, subscribeServerPort, subscribeServiceHealthEndpoint)
)
