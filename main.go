package main

import (
	"net/http"
	"time"

	"github.com/JayDalton/simplegoserver/metrics"
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
	collection := metrics.NewCollection(time.Now())

	router := httprouter.New()

	router.GET(routes.RootName, routes.GetRoot(collection, "Hallo, Welt!"))
	router.GET(routes.HealthName, routes.GetHealth(collection))

	server := http.Server{Addr: ":3000", Handler: router}
	err := server.ListenAndServe()

	log.Fatal().Err(err).Msg("server failed")
}
