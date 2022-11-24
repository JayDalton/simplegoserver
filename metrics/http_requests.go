package metrics

type HttpRequestsType struct {
	Counters map[string]int64
}

var HttpRequests = HttpRequestsType{
	Counters: make(map[string]int64),
}
