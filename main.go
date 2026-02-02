func getFlightsHandler(w http.ResponseWriter, r *http.Request) {
    flightType := r.URL.Query().Get("type")
    if flightType == "" { flightType = "arrivals" }
    
    // Using a slightly different internal URL that is often used by their mobile app
    apiURL := fmt.Sprintf("https://www.dubaiairports.ae/api/flights?type=%s&date=2026-02-02", flightType)

    client := &http.Client{}
    req, _ := http.NewRequest("GET", apiURL, nil)

    // Masking our server as a real Chrome browser
    req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
    req.Header.Set("Accept", "application/json, text/plain, */*")
    req.Header.Set("Accept-Language", "en-US,en;q=0.9")
    req.Header.Set("Referer", "https://www.dubaiairports.ae/flight-status")
    req.Header.Set("X-Requested-With", "XMLHttpRequest")

    resp, err := client.Do(req)
    if err != nil {
        http.Error(w, "Failed to connect to DXB", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    // If we still get HTML, let's log the status code
    if resp.StatusCode != 200 {
        fmt.Printf("Error: Received status code %d\n", resp.StatusCode)
    }

    w.Header().Set("Content-Type", "application/json")
    io.Copy(w, resp.Body)
}
