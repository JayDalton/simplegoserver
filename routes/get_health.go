package routes

import (
	"encoding/json"
	"net/http"

	"github.com/JayDalton/simplegoserver/metrics"
	"github.com/julienschmidt/httprouter"
)

const HealthName = "/health"

func GetHealth(collection *metrics.Collection) RouteFunctor {
	collection.HttpRequests.Counters[HealthName] = 0

	return func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		defer func() {
			collection.HttpRequests.Counters[HealthName]++
		}()

		uptime, unit := collection.Uptime.Value()
		reqCounters := collection.HttpRequests.Counters

		response := HealthResponse{
			Uptime:       UptimeResponse{Value: uptime, Unit: unit},
			HttpRequests: HttpRequestsResponse{Counters: reqCounters},
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
