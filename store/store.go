package store

// Operations is the interface which will be implemented by all the stores
// like inMemory store, db store etc
type Operations interface {
	Add(ip string, metrics *Metrics)
}
