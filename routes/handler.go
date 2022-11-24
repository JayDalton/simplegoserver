package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type RouteFunctor = func(http.ResponseWriter, *http.Request, httprouter.Params)
