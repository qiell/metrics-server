package handler

import (
	"encoding/json"
	"github.com/qiell/metrics-server/constants"
	"github.com/qiell/metrics-server/events"
	"github.com/qiell/metrics-server/store"
	"net/http"
)

// MetricsHandler is responsible to handle the request sent on /metrics route
func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	// if request method is of type POST, then process the request,
	// otherwise return 500
	if r.Method == constants.Post {
		// create object of Metrics
		var metrics store.Metrics

		// decode request body into metrics object
		err := json.NewDecoder(r.Body).Decode(&metrics)
		if err != nil {
			http.Error(w, "Error in parsing request body", http.StatusInternalServerError)
			return
		}

		// add metrics to Store
		store.Object.Add(r.RemoteAddr, &metrics)
		return
	} else {
		http.Error(w, "Invalid request method", http.StatusInternalServerError)
		return
	}
}

// ReportHandler is responsible to handle the request sent on /report route
func ReportHandler(w http.ResponseWriter, r *http.Request) {
	// if request method is of type GET, then process the request,
	// otherwise return 500
	if r.Method == constants.Get {
		// create an array of Metrics object
		metrics := make([]*events.Metrics, 0)
		for _, value := range events.MaxMetrics {
			// append the maximum metrics of each client in metrics object
			metrics = append(metrics, value)
		}

		// Marshal the metrics object to json
		response, err := json.Marshal(metrics)
		if err != nil {
			http.Error(w, "Error in marshalling response", http.StatusInternalServerError)
		}

		// set content type header
		w.Header().Set(constants.ContentTypeHeader, constants.ApplicationJSON)
		// send response
		w.Write(response)
		return
	} else {
		http.Error(w, "Invalid request method", http.StatusInternalServerError)
		return
	}
}
