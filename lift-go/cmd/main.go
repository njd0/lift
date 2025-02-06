package main

import (
	"context"
	"log"
	config "m/lift/config"
	db "m/lift/db"
	exercises "m/lift/internal/api/exercises"
	cors "m/lift/internal/middleware"
	"net/http"
)

const keyServerAddr = "serverAddr"

func main() {
	ctx := context.Background()

	config.LoadConfig("config.local.json")

	db.ConnectDB(ctx)

	mux := http.NewServeMux()
	
	// setup endpoints
	mux.HandleFunc("/v1/exercises", exercises.GetExercises)

	// set up middleware
	handler := cors.CorsMiddleware(mux)

	// setup server on port 3333
	server := &http.Server{
		Addr:    ":3333",
		Handler: handler,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("server failure", err)
	}
}