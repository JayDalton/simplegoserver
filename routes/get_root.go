package routes

import (
	"net/http"

	"github.com/JayDalton/simplegoserver/metrics"
	"github.com/julienschmidt/httprouter"
)

const RootName = "/"

func GetRoot(collection *metrics.Collection, message string) RouteFunctor {
	collection.HttpRequests.RegisterRoute(RootName)

	return func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		defer collection.HttpRequests.Increment(RootName)

		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(message))
	}
}
