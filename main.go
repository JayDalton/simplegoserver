package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
)

func main() {
	fmt.Println("Hallo Welt!")
	router := httprouter.New()

	router.GET("/", getRoot)

	server := http.Server{Addr: ":3000", Handler: router}
	err := server.ListenAndServe()

	log.Fatal().Err(err).Msg("server failed")
}

func getRoot(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Hallo, Welllt!!"))
}
