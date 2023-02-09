package main

import (
	"fmt"
	"time"

	_ "github.com/joho/godotenv/autoload"
	as "github.com/palavrapasse/aspirador/pkg"
	"github.com/palavrapasse/monitor/internal/http"
	"github.com/palavrapasse/monitor/internal/logging"
	"github.com/palavrapasse/paramedic/pkg"
)

const WaitingSecondsBetweenHealthChecks = 600
const MinPercentageToBeOverloaded = 20

var WebServicesHeathURL = map[string]string{http.QueryServiceName: http.QueryServiveHealthURL, http.SantosServiceName: http.SantosServiveHealthURL, http.SubscribeServiceName: http.SubscribeServiveHealthURL}

func main() {

	logging.Aspirador = as.WithClients(logging.CreateAspiradorClients())

	logging.Aspirador.Trace("Starting Monitor Service")

	for {

		for serviceName, url := range WebServicesHeathURL {
			result, err := http.GetServiceHealth(url)

			if err != nil {
				logging.Aspirador.Error(err.Error())
				continue
			}

			if err = validateResult(result); err != nil {
				logging.Aspirador.Warning(fmt.Sprintf("%s %s ", serviceName, err.Error()))
				continue
			}
		}

		logging.Aspirador.Trace(fmt.Sprintf("Waiting %d seconds...", WaitingSecondsBetweenHealthChecks))
		time.Sleep(WaitingSecondsBetweenHealthChecks * time.Second)
	}
}

func validateResult(r pkg.HealthStatus) error {
	var err error

	if err = validatePercentage(r.CPUPercentage); err == nil {
		err = validatePercentage(r.RAMPercentage)
	}

	if err != nil {
		err = fmt.Errorf("%s (CPU percentage is %.2f and RAM percentage is %.2f)", err.Error(), r.CPUPercentage, r.RAMPercentage)
	}

	return err
}

func validatePercentage(p float64) error {
	if p > MinPercentageToBeOverloaded {
		return fmt.Errorf("service is overloaded")
	}

	return nil
}
