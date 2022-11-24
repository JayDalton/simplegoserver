package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
)

type HttpLambda = func(http.ResponseWriter, *http.Request, httprouter.Params)

func main() {
	router := httprouter.New()

	router.GET("/", getRoot("huhu ahoi"))

	server := http.Server{Addr: ":3000", Handler: router}
	err := server.ListenAndServe()

	log.Fatal().Err(err).Msg("server failed")
}

func getRoot(message string) HttpLambda {
	return func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(message))
	}
}
