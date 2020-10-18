package store

// Metrics struct will be used to parse the metrics json response of a client
type Metrics struct {
	CPU int `json:"percentage_cpu_used"`
	Memory int `json:"percentage_memory_used"`
}

// Object is a global variable, it will be used to add the metrics to the store
var Object Operations