package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Flight struct {
	FlightNo     string `json:"flight_no"`
	Airline      string `json:"airline"`
	Origin       string `json:"origin"`
	ScheduleTime string `json:"schedule_time"`
	Status       string `json:"status"`
}

func getFlightsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Наша временная база данных (Mock Data)
	flights := []Flight{
		{
			FlightNo:     "EK132",
			Airline:      "Emirates",
			Origin:       "Moscow (DME)",
			ScheduleTime: "16:45",
			Status:       "On Time",
		},
		{
			FlightNo:     "FZ728",
			Airline:      "flydubai",
			Origin:       "Istanbul (IST)",
			ScheduleTime: "17:10",
			Status:       "Delayed",
		},
	}

	json.NewEncoder(w).Encode(flights)
}

func main() {
	http.HandleFunc("/flights", getFlightsHandler)

	// Берем порт из системы (важно для Render!)
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000" 
	}

	fmt.Printf("Server starting on port %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}