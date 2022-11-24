package main

import (
	"net/http"

	"github.com/JayDalton/simplegoserver/routes"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
)

// go test -cover ./...
// go tooldist list
// go env
// $env:GOOS=abc
// $env:GOARCH=xyz

func main() {
	router := httprouter.New()

	router.GET(routes.RootName, routes.GetRoot("Hallo, Welt!"))
	router.GET(routes.HealthName, routes.GetHealth())

	server := http.Server{Addr: ":3000", Handler: router}
	err := server.ListenAndServe()

	log.Fatal().Err(err).Msg("server failed")
}
