package routes

import (
	"net/http"

	"github.com/JayDalton/simplegoserver/metrics"
	"github.com/julienschmidt/httprouter"
)

const RootName = "/"

func GetRoot(collection *metrics.Collection, message string) RouteFunctor {
	collection.HttpRequests.Counters[RootName] = 0

	return func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		defer func() {
			collection.HttpRequests.Counters[RootName]++
		}()

		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(message))
	}
}
