package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/veanusnathan/rssagg/internal/database"
)

func (apiconfig *apiconfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameter struct {
		Name string `json:"name"`
	}

	param := parameter{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&param)

	if err != nil {
		ErrResponse(w, 400, fmt.Sprintf("Invalid request body: %v", err))

	}

	user, err := apiconfig.Db.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      param.Name,
	})
	if err != nil {
		ErrResponse(w, 500, fmt.Sprintf("Failed to create user: %v", err))
		return
	}

	JSONResponse(w, 200, dbUserToUser(user))
}
