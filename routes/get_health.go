package routes

import (
	"encoding/json"
	"net/http"

	"github.com/JayDalton/simplegoserver/metrics"
	"github.com/julienschmidt/httprouter"
)

const HealthName = "/health"

func GetHealth() RouteFunctor {
	metrics.HttpRequests.Counters[HealthName] = 0

	return func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		metrics.HttpRequests.Counters[HealthName]++

		uptime, unit := metrics.Uptime()
		reqCounters := metrics.HttpRequests.Counters

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
