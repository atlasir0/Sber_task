package main

import (
	"log"
	"net/http"
	"todolist/internal/api"
	"todolist/internal/db"
	"todolist/internal/handlers_api"
	"todolist/internal/repositories"
	"todolist/internal/services"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	db, err := db.InitDB()
	if err != nil {
		log.Fatalf("could not connect to the database: %v", err)
	}

	apiInstance, err := api.NewAPI(router)
	if err != nil {
		log.Fatalf("could not create API instance: %v", err)
	}
	apiInstance.RegisterHandlers()

	taskRepo := repositories.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepo)
	taskHandler := handlers_api.NewTaskHandler(taskService)
	taskHandler.RegisterRoutes(router)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
