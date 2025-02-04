package service

import (
	"github.com/go_web/internal/models"
	"github.com/go_web/internal/repository"
)

type Authorization interface {
	CreateUser(c models.User) error
	GenerateJWTToken(username string, password string) (string, error)
	GetID(username string, password string) (string, error)
	ParsingJWTToken(token string) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
