package models

import "time"

type FlightInformation struct {
	Pagination struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		Count  int `json:"count"`
		Total  int `json:"total"`
	} `json:"pagination"`
	Data []struct {
		FlightDate   string `json:"flight_date"`
		FlightStatus string `json:"flight_status"`
		Departure    struct {
			Airport         string    `json:"airport"`
			Timezone        string    `json:"timezone"`
			Iata            string    `json:"iata"`
			Icao            string    `json:"icao"`
			Terminal        string    `json:"terminal"`
			Gate            string    `json:"gate"`
			Delay           int       `json:"delay"`
			Scheduled       time.Time `json:"scheduled"`
			Estimated       time.Time `json:"estimated"`
			Actual          time.Time `json:"actual"`
			EstimatedRunway time.Time `json:"estimated_runway"`
			ActualRunway    time.Time `json:"actual_runway"`
		} `json:"departure"`
		Arrival struct {
			Airport         string      `json:"airport"`
			Timezone        string      `json:"timezone"`
			Iata            string      `json:"iata"`
			Icao            string      `json:"icao"`
			Terminal        string      `json:"terminal"`
			Gate            string      `json:"gate"`
			Baggage         string      `json:"baggage"`
			Delay           int         `json:"delay"`
			Scheduled       time.Time   `json:"scheduled"`
			Estimated       time.Time   `json:"estimated"`
			Actual          interface{} `json:"actual"`
			EstimatedRunway interface{} `json:"estimated_runway"`
			ActualRunway    interface{} `json:"actual_runway"`
		} `json:"arrival"`
		Airline struct {
			Name string `json:"name"`
			Iata string `json:"iata"`
			Icao string `json:"icao"`
		} `json:"airline"`
		Flight struct {
			Number     string      `json:"number"`
			Iata       string      `json:"iata"`
			Icao       string      `json:"icao"`
			Codeshared interface{} `json:"codeshared"`
		} `json:"flight"`
		Aircraft struct {
			Registration string `json:"registration"`
			Iata         string `json:"iata"`
			Icao         string `json:"icao"`
			Icao24       string `json:"icao24"`
		} `json:"aircraft"`
		Live struct {
			Updated         time.Time `json:"updated"`
			Latitude        float64   `json:"latitude"`
			Longitude       float64   `json:"longitude"`
			Altitude        float64   `json:"altitude"`
			Direction       float64   `json:"direction"`
			SpeedHorizontal float64   `json:"speed_horizontal"`
			SpeedVertical   float64   `json:"speed_vertical"`
			IsGround        bool      `json:"is_ground"`
		} `json:"live"`
	} `json:"data"`
}