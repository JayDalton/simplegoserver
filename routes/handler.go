package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HttpLambda = func(http.ResponseWriter, *http.Request, httprouter.Params)
