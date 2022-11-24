package metrics

type HttpRequests struct {
	Counters map[string]int64
}

func NewHttpRequests() *HttpRequests {
	return &HttpRequests{
		Counters: map[string]int64{},
	}
}
