package repository

import (
	"database/sql"
	"todo-challange/model"
)

// interface
type UserRepository interface {
	GetById(id string) (model.User, error)
}

// struct
type userRepository struct {
	db *sql.DB
}

func (r *userRepository) GetById(id string) (model.User, error) {
	var user model.User

	err := r.db.QueryRow("SELECT id, fullname, email, passwords, role, created_at, updated_at FROM mst_users WHERE id = $1", id).Scan(&user.Id, &user.Fullname, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

// constructor
func NewUserRepository(database *sql.DB) UserRepository {
	return &userRepository{db: database}
}
