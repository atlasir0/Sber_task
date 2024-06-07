package api

import (
	"database/sql"
	"todolist/internal/handlers_api"
	"todolist/internal/repositories"
	"todolist/internal/services"

	"github.com/gorilla/mux"
)

type API struct {
	DB     *sql.DB
	Router *mux.Router
}

func NewAPI(router *mux.Router, dbConn *sql.DB) (*API, error) {
	return &API{
		DB:     dbConn,
		Router: router,
	}, nil
}

func (a *API) RegisterHandlers() {
	taskRepo := repositories.NewTaskRepository(a.DB)
	taskService := services.NewTaskService(taskRepo)
	taskHandler := handlers_api.NewTaskHandler(taskService)

	taskHandler.RegisterRoutes(a.Router)
}
