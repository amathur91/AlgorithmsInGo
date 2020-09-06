package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"server/flightservice/models"
)

func GetDataFromExternalService() models.FlightInformation {
	url := "http://api.aviationstack.com/v1/flights?access_key=aa2629845870ca0f87a711d311b8bc72"
	resp, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close()
	var flightInformation models.FlightInformation
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(body, &flightInformation)
	return flightInformation
}
