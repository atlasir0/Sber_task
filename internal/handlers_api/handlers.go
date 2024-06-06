package handlers_api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"todolist/internal/models"
	"todolist/internal/services"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
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

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.TaskService.CreateTask(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.TaskService.ListTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task, err := h.TaskService.GetTaskByID(taskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var task models.Task
	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	task.ID = taskID

	err = h.TaskService.UpdateTask(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	err = h.TaskService.DeleteTask(taskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
func (h *TaskHandler) FilterTasksByDateAndStatus(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Inside FilterTasksByDateAndStatus function")

	dateStr := r.URL.Query().Get("date")
	statusStr := r.URL.Query().Get("status")

	logrus.Printf("Date: %s, Status: %s\n", dateStr, statusStr)

	if dateStr == "" || statusStr == "" {
		http.Error(w, "Missing date or status parameter", http.StatusBadRequest)
		return
	}

	date, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		logrus.Errorf("Error parsing date: %v", err)
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	completed, err := strconv.ParseBool(statusStr)
	if err != nil {
		logrus.Errorf("Error parsing status: %v", err)
		http.Error(w, "Invalid status format", http.StatusBadRequest)
		return
	}

	tasks, err := h.TaskService.FilterTasksByDateAndStatus(date, completed)
	if err != nil {
		logrus.Errorf("Error filtering tasks: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) ListPaginatedTasks(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	statusStr := r.URL.Query().Get("status")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		http.Error(w, "Invalid page parameter", http.StatusBadRequest)
		logrus.Errorf("Invalid page parameter: %v", err)
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
		logrus.Errorf("Invalid limit parameter: %v", err)
		return
	}
	status, err := strconv.ParseBool(statusStr)
	if err != nil {
		http.Error(w, "Invalid status parameter", http.StatusBadRequest)
		logrus.Errorf("Invalid status parameter: %v", err)
		return
	}

	tasks, err := h.TaskService.ListPaginatedTasks(page, limit, status)
	if err != nil {
		http.Error(w, "Failed to get tasks", http.StatusInternalServerError)
		logrus.Errorf("Failed to get tasks: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
