package routes

import (
	"encoding/json"
	"net/http"

	"github.com/JayDalton/simplegoserver/metrics"
	"github.com/julienschmidt/httprouter"
)

type HealthResponse struct {
	Uptime UptimeResponse `json:"uptime"`
}

type UptimeResponse struct {
	Value int64  `json:"value"`
	Unit  string `json:"unit"`
}

func GetHealth() RouteFunctor {
	return func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		uptime, unit := metrics.Uptime()

		response := HealthResponse{
			Uptime: UptimeResponse{Value: uptime, Unit: unit},
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
