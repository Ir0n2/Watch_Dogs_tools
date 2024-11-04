package main

import (
    "fmt"
    "net/http"
    "log"
)

func main() {
    http.HandleFunc("/", serveHomePage)
    http.HandleFunc("/location", handleLocation)

    fmt.Println("Starting server on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

// Serve the HTML page
func serveHomePage(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")
}

// Handle the location data sent from the client
func handleLocation(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        latitude := r.FormValue("latitude")
        longitude := r.FormValue("longitude")

        fmt.Println("Received location data:")
        fmt.Println("Latitude:", latitude)
        fmt.Println("Longitude:", longitude)

        w.Write([]byte("Location received successfully!"))
    } else {
        http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
    }
}

