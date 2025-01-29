package repository

import (
	"database/sql"
	"github.com/go_web/internal/models"
)

type Authorization interface {
	CreateUser(client models.User) error
	GetUserID(username string, password string) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
