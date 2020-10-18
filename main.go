package main

import (
	"github.com/qiell/metrics-server/constants"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/metrics", metricsHandler)
	http.HandleFunc("/report", reportHandler)

	log.Println("Starting server at http://localhost" + constants.Port)
	http.ListenAndServe(constants.Port, nil)
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == constants.Post {
		w.Write([]byte("request successful"))
		return
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
}

func reportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == constants.Get {
		w.Write([]byte("request successful"))
		return
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
}