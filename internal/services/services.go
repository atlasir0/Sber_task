package services

import (
	"fmt"
	"time"
	"todolist/internal/models"
	"todolist/internal/repositories"
)

type TaskService struct {
	TaskRepo *repositories.TaskRepository
}

func NewTaskService(taskRepo *repositories.TaskRepository) *TaskService {
	return &TaskService{
		TaskRepo: taskRepo,
	}
}

func (s *TaskService) CreateTask(task *models.Task) error {
	return s.TaskRepo.Create(task)
}

func (s *TaskService) ListTasks() ([]models.Task, error) {
	return s.TaskRepo.GetAll()
}

func (s *TaskService) GetTaskByID(id int) (models.Task, error) {
	return s.TaskRepo.GetByID(id)
}

func (s *TaskService) UpdateTask(task *models.Task) error {
	return s.TaskRepo.Update(task)
}

func (s *TaskService) DeleteTask(id int) error {
	return s.TaskRepo.Delete(id)
}

func (s *TaskService) FilterTasksByDateAndStatus(date time.Time, completed bool) ([]models.Task, error) {
	fmt.Print("S")
	return s.TaskRepo.GetTasksByDateAndStatus(date, completed)
}

func (s *TaskService) ListPaginatedTasks(page, limit int, status bool) ([]models.Task, error) {

	offset := (page - 1) * limit

	return s.TaskRepo.GetPaginatedTasks(status, offset, limit)
}
