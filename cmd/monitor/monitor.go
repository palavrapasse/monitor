package main

import (
	"fmt"
	"time"

	_ "github.com/joho/godotenv/autoload"
	as "github.com/palavrapasse/aspirador/pkg"
	"github.com/palavrapasse/monitor/internal/http"
	"github.com/palavrapasse/monitor/internal/logging"
)

const WaitingSecondsBetweenHealthChecks = 600

var WebServicesHeathURL = []string{http.QueryServiveHealthURL, http.SantosServiveHealthURL, http.SubscribeServiveHealthURL}

func main() {

	logging.Aspirador = as.WithClients(logging.CreateAspiradorClients())

	logging.Aspirador.Trace("Starting Monitor Service")

	for {

		for _, url := range WebServicesHeathURL {
			_, err := http.GetServiceHealth(url)

			if err != nil {
				logging.Aspirador.Error(err.Error())
			}
		}

		logging.Aspirador.Trace(fmt.Sprintf("Waiting %d seconds...", WaitingSecondsBetweenHealthChecks))
		time.Sleep(WaitingSecondsBetweenHealthChecks * time.Second)
	}
}
