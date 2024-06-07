package handlers_api

import (
	"todolist/internal/services"

	"github.com/gorilla/mux"
)


func (h *TaskHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/tasks/paginated", h.ListPaginatedTasks).Methods("GET")
	router.HandleFunc("/tasks", h.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/filter", h.FilterTasksByDateAndStatus).Methods("GET")
	router.HandleFunc("/tasks", h.ListTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", h.GetTaskByID).Methods("GET")
	router.HandleFunc("/tasks/{id}", h.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", h.DeleteTask).Methods("DELETE")
}

type TaskHandler struct {
	TaskService *services.TaskService
}


func NewTaskHandler(taskService *services.TaskService) *TaskHandler {
	return &TaskHandler{
		TaskService: taskService,
	}
}
