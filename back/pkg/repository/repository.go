package repository

import (
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/models"
	"github.com/jmoiron/sqlx"
)

type Authorisation interface {
	CreateUser(user models.User) (int, error)
	GetUser(email string, passwordHash string) (models.User, error)
	GetUserById(userId int) (models.User, error)
}

type Purchase interface {
	AddPurchase(purchase models.Purchase) (int, error)
}

type Repository struct {
	Authorisation
	Purchase
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorisation: NewAuth(db),
		Purchase:      NewPurchase(db),
	}
}
