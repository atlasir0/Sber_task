-- name: CreateTask :one
INSERT INTO tasks (title, description, date, completed) 
VALUES ($1, $2, $3, $4) 
RETURNING id, title, description, date, completed;

-- name: GetAllTasks :many
SELECT id, title, description, date, completed FROM tasks;

-- name: GetTaskByID :one
SELECT id, title, description, date, completed 
FROM tasks 
WHERE id = $1;

-- name: UpdateTask :exec
UPDATE tasks 
SET title = $1, description = $2, date = $3, completed = $4 
WHERE id = $5;

-- name: DeleteTask :exec
DELETE FROM tasks 
WHERE id = $1;
