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

func main() {

	logging.Aspirador = as.WithClients(logging.CreateAspiradorClients())

	logging.Aspirador.Trace("Starting Monitor Service")

	for {
		_, santosErr := http.GetSantosHealth()
		if santosErr != nil {
			logging.Aspirador.Error(santosErr.Error())
		}

		_, queryErr := http.GetQueryHealth()
		if queryErr != nil {
			logging.Aspirador.Error(queryErr.Error())
		}

		_, subscribeErr := http.GetSubscribeHealth()
		if subscribeErr != nil {
			logging.Aspirador.Error(subscribeErr.Error())
		}

		logging.Aspirador.Trace(fmt.Sprintf("Waiting %d seconds...", WaitingSecondsBetweenHealthChecks))
		time.Sleep(WaitingSecondsBetweenHealthChecks * time.Second)
	}
}
