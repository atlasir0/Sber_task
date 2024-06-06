package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"todolist/internal/db"
	"todolist/internal/models"

	log "github.com/sirupsen/logrus"
)

type TaskRepository struct {
	Queries *db.Queries
}

func NewTaskRepository(dbConn *sql.DB) *TaskRepository {
	return &TaskRepository{
		Queries: db.New(dbConn),
	}
}

func (r *TaskRepository) Create(task *models.Task) error {
	createdTask, err := r.Queries.CreateTask(context.Background(), db.CreateTaskParams{
		Title:       task.Title,
		Description: sql.NullString{String: task.Description, Valid: true},
		Date:        task.Date,
		Completed:   task.Completed,
	})
	if err != nil {
		log.Errorf("failed to create task: %v", err)
		return err
	}
	task.ID = int(createdTask.ID)
	return nil
}

func (r *TaskRepository) GetAll() ([]models.Task, error) {
	tasks, err := r.Queries.GetAllTasks(context.Background())
	if err != nil {
		log.Errorf("failed to get all tasks: %v", err)
		return nil, err
	}

	var result []models.Task
	for _, task := range tasks {
		result = append(result, models.Task{
			ID:          int(task.ID),
			Title:       task.Title,
			Description: task.Description.String,
			Date:        task.Date,
			Completed:   task.Completed,
		})
	}
	return result, nil
}

func (r *TaskRepository) GetByID(id int) (models.Task, error) {
	task, err := r.Queries.GetTaskByID(context.Background(), int32(id))
	if err != nil {
		log.Errorf("failed to get task by ID: %v", err)
		return models.Task{}, err
	}
	return models.Task{
		ID:          int(task.ID),
		Title:       task.Title,
		Description: task.Description.String,
		Date:        task.Date,
		Completed:   task.Completed,
	}, nil
}

func (r *TaskRepository) Update(task *models.Task) error {
	err := r.Queries.UpdateTask(context.Background(), db.UpdateTaskParams{
		Title:       task.Title,
		Description: sql.NullString{String: task.Description, Valid: true},
		Date:        task.Date,
		Completed:   task.Completed,
		ID:          int32(task.ID),
	})
	if err != nil {
		log.Errorf("failed to update task: %v", err)
	}
	return err
}

func (r *TaskRepository) Delete(id int) error {
	err := r.Queries.DeleteTask(context.Background(), int32(id))
	if err != nil {
		log.Errorf("failed to delete task: %v", err)
	}
	return err
}

func (r *TaskRepository) GetTasksByDateAndStatus(date time.Time, completed bool) ([]models.Task, error) {
	fmt.Print("S")
	tasks, err := r.Queries.GetTasksByDateAndStatus(context.Background(), db.GetTasksByDateAndStatusParams{
		Date:      date,
		Completed: completed,
	})
	if err != nil {
		log.Errorf("failed to get tasks by date and status: %v", err)
		return nil, err
	}

	var result []models.Task
	for _, task := range tasks {
		result = append(result, models.Task{
			ID:          int(task.ID),
			Title:       task.Title,
			Description: task.Description.String,
			Date:        task.Date,
			Completed:   task.Completed,
		})
	}
	return result, nil
}

func (r *TaskRepository) GetPaginatedTasks(completed bool, offset, limit int) ([]models.Task, error) {
	tasks, err := r.Queries.GetPaginatedTasks(context.Background(), db.GetPaginatedTasksParams{
		Completed: completed,
		Offset:    int32(offset),
		Limit:     int32(limit),
	})
	if err != nil {
		return nil, err
	}

	var result []models.Task
	for _, task := range tasks {
		result = append(result, models.Task{
			ID:          int(task.ID),
			Title:       task.Title,
			Description: task.Description.String,
			Date:        task.Date,
			Completed:   task.Completed,
		})
	}
	return result, nil
}
