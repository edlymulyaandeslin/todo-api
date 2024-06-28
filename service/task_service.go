package service

import (
	"todo-challange/model"
	"todo-challange/model/dto"
	"todo-challange/repository"
)

type TaskService interface {
	FindById(id string) (model.Task, error)
	FindAllTask() ([]model.Task, error)
	CreateNewTask(payload dto.TaskRequest) (model.Task, error)
	UpdatedTask(id string, payload dto.TaskUpdated) (model.Task, error)
	DeleteTask(id string) error
}

type taskService struct {
	repo repository.TaskRepository
	uS   UserService
}

func (s *taskService) FindById(id string) (model.Task, error) {
	task, err := s.repo.GetById(id)
	if err != nil {
		return model.Task{}, err
	}

	user, err := s.uS.FindById(task.User.Id)
	if err != nil {
		return model.Task{}, err
	}

	task.User = user
	return task, nil
}

func (s *taskService) FindAllTask() ([]model.Task, error) {

	var listDataTask []model.Task
	dataTask, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	for _, task := range dataTask {
		user, err := s.uS.FindById(task.User.Id)
		if err != nil {
			return nil, err
		}
		task.User = user
		listDataTask = append(listDataTask, task)
	}

	return listDataTask, nil
}

func (s *taskService) CreateNewTask(payload dto.TaskRequest) (model.Task, error) {
	user, err := s.uS.FindById(payload.UserId)
	if err != nil {
		return model.Task{}, err
	}

	newPayload := model.Task{
		Title:   payload.Title,
		Content: payload.Content,
		User:    user,
	}

	task, err := s.repo.CreateTask(newPayload)
	if err != nil {
		return model.Task{}, err
	}

	task.User = user
	return task, nil
}

func (s *taskService) UpdatedTask(id string, payload dto.TaskUpdated) (model.Task, error) {

	task, err := s.repo.GetById(id)
	if err != nil {
		return model.Task{}, err
	}

	user, err := s.uS.FindById(task.User.Id)
	if err != nil {
		return model.Task{}, err
	}

	if payload.Title == "" {
		payload.Title = task.Title
	}
	if payload.Content == "" {
		payload.Content = task.Content
	}

	newPayloadUpdate := model.Task{
		Title:   payload.Title,
		Content: payload.Content,
	}

	taskUpdate, err := s.repo.UpdateTask(id, newPayloadUpdate)
	if err != nil {
		return model.Task{}, err
	}

	taskUpdate.User = user
	taskUpdate.CreatedAt = task.CreatedAt

	return taskUpdate, nil
}

func (s *taskService) DeleteTask(id string) error {
	return s.repo.Delete(id)
}

// constructor
func NewTaskService(repo repository.TaskRepository, uS UserService) TaskService {
	return &taskService{repo: repo, uS: uS}
}
