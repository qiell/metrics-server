package events

// NewBroker will instantiate the broker object
func NewBroker() (broker *Broker) {
	// instantiate the broker
	broker = &Broker{
		Notifier: make(chan Metrics),
	}

	// set it running
	go broker.listen()
	return
}

// listen on notifier channel and act accordingly
func (b *Broker) listen() {
	for {
		select {
		case metrics := <-b.Notifier:
			// calculate the max metrics and store in it MaxMetrics object
			metrics.CalculateMax()
		}
	}
}
