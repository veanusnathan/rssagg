package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func ErrResponse(w http.ResponseWriter, code int, msg string) {
	if code >= 500 {
		log.Println("Responding with 5xx error : ", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	JSONResponse(w, code, errResponse{
		Error: msg,
	})

}

func JSONResponse(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON Response %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
