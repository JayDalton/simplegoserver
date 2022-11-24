package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetRoot(message string) HttpLambda {
	return func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(message))
	}
}
