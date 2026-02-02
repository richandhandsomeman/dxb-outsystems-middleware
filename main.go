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

	// 2. Target URL
	apiURL := fmt.Sprintf("https://www.dubaiairports.ae/api/flights?type=%s&date=2026-02-02", flightType)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", apiURL, nil)

	// 3. Browser emulation headers
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", "https://www.dubaiairports.ae/flight-status")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to connect to DXB", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// 4. Send response back to OutSystems
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
