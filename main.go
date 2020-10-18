package main

import (
	"github.com/qiell/metrics-server/constants"
	"log"
	"net/http"
)

func main() {
	// add metrics route
	http.HandleFunc(constants.Metrics, metricsHandler)
	// add report route
	http.HandleFunc(constants.Report, reportHandler)

	// starting server
	log.Println("Starting server at http://localhost" + constants.Port)
	http.ListenAndServe(constants.Port, nil)
}

// metricsHandler is responsible to handle the request sent on /metrics route
func metricsHandler(w http.ResponseWriter, r *http.Request) {
	// if request method is of type POST, then process the request,
	// otherwise return 500
	if r.Method == constants.Post {
		w.Write([]byte("request successful"))
		return
	} else {
		http.Error(w, "Invalid request method", http.StatusInternalServerError)
		return
	}
}

// reportHandler is responsible to handle the request sent on /report route
func reportHandler(w http.ResponseWriter, r *http.Request) {
	// if request method is of type GET, then process the request,
	// otherwise return 500
	if r.Method == constants.Get {
		w.Write([]byte("request successful"))
		return
	} else {
		http.Error(w, "Invalid request method", http.StatusInternalServerError)
		return
	}
}