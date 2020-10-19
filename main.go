package main

import (
	"github.com/qiell/metrics-server/constants"
	"github.com/qiell/metrics-server/events"
	"github.com/qiell/metrics-server/handler"
	"github.com/qiell/metrics-server/store"
	"github.com/qiell/metrics-server/store/inmemory"
	"log"
	"net/http"
)

func init() {
	// TODO: use config to determine the kind of Store like inMemory or DB
	store.Object = inmemory.New()
	events.MaxMetrics = events.NewMaxMetrics()
	events.BrokerObject = events.NewBroker()
}

func main() {
	// add metrics route
	http.HandleFunc(constants.Metrics, handler.MetricsHandler)
	// add report route
	http.HandleFunc(constants.Report, handler.ReportHandler)

	// starting server
	log.Println("Starting server at http://localhost" + constants.Port)
	http.ListenAndServe(constants.Port, nil)
}
