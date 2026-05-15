package main

import (
	"fmt"
	"net/http"
	"time"
)

func eventsHandler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers if needed (for simplicity, we allow all here)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	fmt.Println("Client connected")

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			fmt.Fprintf(w, "data: The time is %s\n\n", t.Format(time.RFC3339))
			flusher.Flush()
		case <-r.Context().Done():
			fmt.Println("Client disconnected")
			return
		}
	}
}

func main() {
	http.HandleFunc("/events", eventsHandler)
	http.Handle("/", http.FileServer(http.Dir(".")))

	fmt.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
