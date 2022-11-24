package routes

import (
	"net/http"

	"github.com/JayDalton/simplegoserver/metrics"
	"github.com/julienschmidt/httprouter"
)

const RootName = "/"

func GetRoot(message string) RouteFunctor {
	metrics.HttpRequests.Counters[RootName] = 0

	return func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		metrics.HttpRequests.Counters[RootName]++

		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(message))
	}
}
