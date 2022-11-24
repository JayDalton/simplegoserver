package metrics

type HttpRequests struct {
	counters map[string]int64
}

func NewHttpRequests() *HttpRequests {
	return &HttpRequests{
		counters: map[string]int64{},
	}
}

func (requests HttpRequests) RegisterRoute(route string) {
	if _, ok := requests.counters[route]; ok {
		panic("route already registered")
	}
	requests.counters[route] = 0
}

func (requests HttpRequests) Increment(route string) {
	if _, ok := requests.counters[route]; !ok {
		panic("route not registered")
	}

	requests.counters[route]++
}

func (requests HttpRequests) Value() map[string]int64 {
	return requests.counters
}
