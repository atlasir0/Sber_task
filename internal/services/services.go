package services

import (
	"todolist/internal/models"
	"todolist/internal/repositories"
)

type TaskService struct {
	taskRepo *repositories.TaskRepository
}

func NewTaskService(taskRepo *repositories.TaskRepository) *TaskService {
	return &TaskService{
		taskRepo: taskRepo,
	}
}

func (s *TaskService) CreateTask(task *models.Task) error {
	return s.taskRepo.Create(task)
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	return s.taskRepo.GetAll()
}

func (s *TaskService) ListTasks() ([]models.Task, error) {
	return s.GetAllTasks()
}

func (s *TaskService) GetTaskByID(id int) (models.Task, error) {
	return s.taskRepo.GetByID(id)
}

func (s *TaskService) UpdateTask(task *models.Task) error {
	return s.taskRepo.Update(task)
}

func (s *TaskService) DeleteTask(id int) error {
	return s.taskRepo.Delete(id)
}
