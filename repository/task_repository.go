package repository

import (
	"database/sql"
	"errors"
	"time"
	"todo-challange/model"
)

type TaskRepository interface {
	GetById(id string) (model.Task, error)
	GetAll() ([]model.Task, error)
	CreateTask(payload model.Task) (model.Task, error)
	UpdateTask(id string, payload model.Task) (model.Task, error)
	Delete(id string) error
}

type taskRepository struct {
	db *sql.DB
}

func (r *taskRepository) GetById(id string) (model.Task, error) {
	var task model.Task

	err := r.db.QueryRow("SELECT id, title, content, user_id, created_at, updated_at FROM trx_tasks WHERE id = $1", id).Scan(&task.Id, &task.Title, &task.Content, &task.User.Id, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func (r *taskRepository) GetAll() ([]model.Task, error) {
	var listTask []model.Task

	rows, err := r.db.Query("SELECT id, title, content, user_id, created_at, updated_at FROM trx_tasks")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task model.Task
		err = rows.Scan(&task.Id, &task.Title, &task.Content, &task.User.Id, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}

		listTask = append(listTask, task)
	}

	return listTask, nil
}

func (r *taskRepository) CreateTask(payload model.Task) (model.Task, error) {
	var task model.Task
	err := r.db.QueryRow("INSERT INTO trx_tasks (title, content, user_id) VALUES($1, $2, $3) RETURNING id, title, content, created_at, updated_at", payload.Title, payload.Content, payload.User.Id).Scan(&task.Id, &task.Title, &task.Content, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func (r *taskRepository) UpdateTask(id string, payload model.Task) (model.Task, error) {

	var task model.Task
	err := r.db.QueryRow("UPDATE trx_tasks SET title = $1, content = $2, updated_at = $3 WHERE id = $4 RETURNING id, title, content, updated_at", payload.Title, payload.Content, time.Now(), id).Scan(&task.Id, &task.Title, &task.Content, &task.UpdatedAt)
	if err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func (r *taskRepository) Delete(id string) error {
	result, err := r.db.Exec("DELETE FROM trx_tasks WHERE id = $1", id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("task not found")
	}

	return nil
}

// constructor
func NewTaskRepository(database *sql.DB) TaskRepository {
	return &taskRepository{db: database}
}
