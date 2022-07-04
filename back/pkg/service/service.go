package service

import (
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/models"
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/repository"
)

type Authorisation interface {
	CreateUser(user models.User) (int, error)
}

type Service struct {
	Authorisation
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Authorisation: NewAuthService(r.Authorisation),
	}
}
