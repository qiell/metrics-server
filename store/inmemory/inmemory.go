package inmemory

import (
	"github.com/qiell/metrics-server/events"
	"github.com/qiell/metrics-server/store"
)

// InMemory struct to store the metrics data in memory
type InMemory struct {
	Metrics map[string][]*store.Metrics
}

// New will instantiate a new object of InMemory struct
func New() *InMemory {
	return &InMemory{
		Metrics: make(map[string][]*store.Metrics),
	}
}

// Add method will append the metrics to the respected ip
func (i *InMemory) Add(ip string, metrics *store.Metrics) {
	i.Metrics[ip] = append(i.Metrics[ip], metrics)
	// create an object of events Metrics
	pushMetrics := events.NewMetrics(ip, metrics.CPU, metrics.Memory)
	// push it to notifier channel
	events.BrokerObject.Notifier <- *pushMetrics
}
