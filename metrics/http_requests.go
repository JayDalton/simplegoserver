package metrics

import "sync"

type HttpRequests struct {
	mutex    sync.RWMutex
	counters map[string]int64
}

func NewHttpRequests() *HttpRequests {
	return &HttpRequests{
		counters: map[string]int64{},
	}
}

func (requests *HttpRequests) RegisterRoute(route string) {
	requests.mutex.Lock()
	defer requests.mutex.Unlock()

	if _, ok := requests.counters[route]; ok {
		panic("route already registered")
	}
	requests.counters[route] = 0
}

func (requests *HttpRequests) Increment(route string) {
	requests.mutex.Lock()
	defer requests.mutex.Unlock()

	if _, ok := requests.counters[route]; !ok {
		panic("route not registered")
	}

	requests.counters[route]++
}

func (requests *HttpRequests) Value() map[string]int64 {
	requests.mutex.RLock()
	defer requests.mutex.RUnlock()

	return requests.counters
}
