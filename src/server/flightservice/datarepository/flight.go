package datarepository
import (
	"database/sql"
	"server/flightservice/models"
)

type DataRepository struct {
	DatabaseHandler *sql.DB
}

func (repo DataRepository) GetDataFromFlight() []models.FlightInformation {
	results, err := repo.DatabaseHandler.Query("select * from flight")
	if err != nil {
		panic(err.Error())
	}
	var flightData []models.FlightInformation
	for results.Next() {
		err = results.Scan(flightData)
		if err != nil {
			panic(err.Error())
		}
	}
	return flightData
}
