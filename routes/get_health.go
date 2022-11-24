package routes

import (
	"encoding/json"
	"net/http"

	"github.com/JayDalton/simplegoserver/metrics"
	"github.com/julienschmidt/httprouter"
)

const HealthName = "/health"

func GetHealth(collection *metrics.Collection) RouteFunctor {
	collection.HttpRequests.RegisterRoute(HealthName)

	return func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		defer collection.HttpRequests.Increment(HealthName)

		uptime, unit := collection.Uptime.Value()
		httpRequests := collection.HttpRequests.Value()

		response := HealthResponse{
			Uptime:       UptimeResponse{Value: uptime, Unit: unit},
			HttpRequests: HttpRequestsResponse{Counters: httpRequests},
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		writer.Write(jsonResponse)
	}
}
