package events

func NewMetrics(ip string, cpu, memory int) *Metrics {
	return &Metrics{
		IP:     ip,
		CPU:    cpu,
		Memory: memory,
	}
}

func NewMaxMetrics() map[string]Metrics {
	metrics := make(map[string]Metrics)
	return metrics
}

func (m *Metrics) CalculateMax() {
	// find the client IP is present in MaxMetrics or not
	metrics, ok := MaxMetrics[m.IP]
	// if IP doesn't exists, then create a key and store the current metrics
	if !ok {
		MaxMetrics[m.IP] = *m
		return
	}

	// if IP is present, then compare the stored cpu and memory with current metrics
	// if current cpu is greater than stored CPU, then update it
	if metrics.CPU < m.CPU {
		metrics.CPU = m.CPU
	}

	// if current memory is greater than stored memory, then update it
	if metrics.Memory < m.Memory {
		metrics.Memory = m.Memory
	}

	// Update the metrics in MaxMetrics object
	MaxMetrics[m.IP] = metrics
}
