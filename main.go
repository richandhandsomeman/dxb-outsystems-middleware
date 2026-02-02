package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func getFlightsHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Get flight type from query parameter (arrivals/departures)
	flightType := r.URL.Query().Get("type")
	if flightType == "" {
		flightType = "arrivals"
	}

	// 2. Get date from query parameter or use today
	date := r.URL.Query().Get("date")
	if date == "" {
		date = "2026-02-02" // You can use time.Now().Format("2006-01-02") later
	}

	apiURL := fmt.Sprintf("https://www.dubaiairports.ae/api/flights?type=%s&date=%s", flightType, date)

	// 3. Create request with headers to bypass basic security
	client := &http.Client{}
	req, _ := http.NewRequest("GET", apiURL, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	req.Header.Set("Referer", "https://www.dubaiairports.ae/flight-status")

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to fetch DXB data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// 4. Proxy the JSON response directly to OutSystems
	w.Header().Set("Content-Type", "application/json")
	io.Copy(w, resp.Body)
}

func main() {
	http.HandleFunc("/flights", getFlightsHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}

	fmt.Printf("Server starting on port %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}
