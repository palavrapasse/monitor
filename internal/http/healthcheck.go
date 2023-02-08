package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/palavrapasse/monitor/internal/logging"
	"github.com/palavrapasse/paramedic/pkg"
)

const URL_FORMAT = "%s:%s/%s"

func GetServiceHealth(url string) (pkg.HealthStatus, error) {
	response, err := httpGetCallService(url)

	if err != nil {
		return pkg.HealthStatus{}, err
	}

	return toHealthStatus(response)
}

func httpGetCallService(url string) (*http.Response, error) {

	logging.Aspirador.Trace(fmt.Sprintf("Calling web service: %s", url))

	resp, err := http.Get(url)

	if err != nil {
		logging.Aspirador.Error(fmt.Sprintf("Error while calling web Service: %s", err))
		return nil, err
	}

	responseStatus := resp.StatusCode
	logging.Aspirador.Trace(fmt.Sprintf("Web service response status: %d", responseStatus))

	if responseStatus != http.StatusOK {
		err = fmt.Errorf("received status code %d from service when calling %s", responseStatus, url)
	}

	return resp, err
}

func toHealthStatus(resp *http.Response) (pkg.HealthStatus, error) {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		logging.Aspirador.Error(fmt.Sprintf("Error while reading body of web service response: %s", err))
		return pkg.HealthStatus{}, err
	}

	var data pkg.HealthStatus

	err = json.Unmarshal(body, &data)
	if err != nil {
		logging.Aspirador.Error(fmt.Sprintf("Error while unmarshal body: %s", err))
		return pkg.HealthStatus{}, err
	}

	return data, err
}
