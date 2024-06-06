package main

import (
	"net/http"
	"todolist/internal/api"
	"todolist/internal/db"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	router := mux.NewRouter()

	_, err := db.InitDB()
	if err != nil {
		log.Fatalf("could not connect to the database: %v", err)
	}

	apiInstance, err := api.NewAPI(router)
	if err != nil {
		log.Fatalf("could not create API instance: %v", err)
	}
	apiInstance.RegisterHandlers()

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
