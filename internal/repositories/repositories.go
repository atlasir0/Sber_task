package repositories

import (
	"context"
	"database/sql"
	"todolist/internal/db"
	"todolist/internal/models"
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
		return err
	}
	task.ID = int(createdTask.ID)
	return nil
}

func (r *TaskRepository) GetAll() ([]models.Task, error) {
	tasks, err := r.Queries.GetAllTasks(context.Background())
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

func (r *TaskRepository) GetByID(id int) (models.Task, error) {
	task, err := r.Queries.GetTaskByID(context.Background(), int32(id))
	if err != nil {
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
	return r.Queries.UpdateTask(context.Background(), db.UpdateTaskParams{
		Title:       task.Title,
		Description: sql.NullString{String: task.Description, Valid: true},
		Date:        task.Date,
		Completed:   task.Completed,
		ID:          int32(task.ID),
	})
}

func (r *TaskRepository) Delete(id int) error {
	return r.Queries.DeleteTask(context.Background(), int32(id))
}
