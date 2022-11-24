package routes

import (
	"net/http"

	"github.com/JayDalton/simplegoserver/metrics"
	"github.com/julienschmidt/httprouter"
)

func GetHealth() HttpLambda {
	return func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(metrics.Uptime().String()))
	}
}
