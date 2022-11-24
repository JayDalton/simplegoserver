package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/JayDalton/simplegoserver/metrics"
	"github.com/julienschmidt/httprouter"
)

type HealthResponse struct {
	Uptime time.Duration `json:"uptime"`
}

func GetHealth() HttpLambda {
	return func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		response := HealthResponse{
			Uptime: metrics.Uptime(),
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		writer.Write(jsonResponse)
		// writer.Write([]byte(metrics.Uptime().String()))
	}
}
