package handlers

import (
	"log"
	"net/http"
	"server/flightservice/datarepository"
)

type Handlers struct{
	DataRepo datarepository.DataRepository
}

func (handle Handlers) GetAllFlights(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching all Flights Data from DataBase")
	result := handle.DataRepo.GetDataFromFlight()
	log.Println(result)
}

func (handle Handlers) SyncFlights(w http.ResponseWriter, r *http.Request) {
	log.Println("Syncing from the External Service")
	log.Println(GetDataFromExternalService())
}