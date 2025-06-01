package main

import (
	"net/http"
)

func handlerHC(w http.ResponseWriter, r *http.Request) {
	JSONResponse(w, 200, struct{}{})
}
