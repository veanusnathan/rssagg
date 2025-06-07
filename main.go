package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/veanusnathan/rssagg/internal/database"

	_ "github.com/lib/pq"
)

type apiconfig struct {
	Db *database.Queries
}

func main() {

	godotenv.Load(".env")
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT not defined")
	}

	dbString := os.Getenv("db_url")
	if dbString == "" {
		log.Fatal("dbString not defined")
	}

	conn, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	queries := database.New(conn)

	apiCfg := apiconfig{
		Db: queries,
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	router.Mount("/v1", v1Router)

	v1Router.Get("/healthz", handlerHC)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", apiCfg.handlerCreateUser)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("RSS Aggregator server starting at port %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Failed starting server, Error :", err)
	}
}
