package events

type Metrics struct {
	IP     string `json:"ip"`
	CPU    int    `json:"cpu"`
	Memory int    `json:"memory"`
}

type Broker struct {
	Notifier chan Metrics
}

var MaxMetrics map[string]*Metrics

var BrokerObject *Broker
