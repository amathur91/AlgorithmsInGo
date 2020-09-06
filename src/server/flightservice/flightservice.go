package flightservice

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	. "server/flightservice/datarepository"
	. "server/flightservice/handlers"
)

var DataRepo DataRepository
var Handle Handlers

func BuildService() {
	muxHandler := setupMux()
	initializeDependencies()
	registerHandlers(muxHandler)
	startService(muxHandler)
}

func initializeDependencies() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/aviation")
	if err != nil {
		panic(err.Error())
	}
	DataRepo.DatabaseHandler = db
	//defer db.Close()
}

func startService(muxHandler *mux.Router) {
	log.Fatal(http.ListenAndServe(":8099", muxHandler))
}

func registerHandlers(muxHandler *mux.Router) {
	Handle = Handlers{DataRepo: DataRepo}
	muxHandler.HandleFunc("/flights", Handle.GetAllFlights)
	muxHandler.HandleFunc("/sync", Handle.SyncFlights)
}

func setupMux() *mux.Router {
	muxHandler := mux.NewRouter()
	return muxHandler
}
