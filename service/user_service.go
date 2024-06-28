package service

import (
	"todo-challange/model"
	"todo-challange/repository"
)

// interface
type UserService interface {
	FindById(id string) (model.User, error)
}

// struct
type userService struct {
	repo repository.UserRepository
}

func (s *userService) FindById(id string) (model.User, error) {
	user, err := s.repo.GetById(id)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

// constructor
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}
