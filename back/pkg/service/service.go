package service

import (
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/models"
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/repository"
)

type Authorisation interface {
	CreateUser(user models.User) (int, error)
	GetUser(email string, password string) (*models.User, error)
	GetUserById(userId int) (models.User, error)
	GenerateToken(user models.User) (string, error)
	ParseToken(s string) (int, error)
}

type Purchase interface {
	AddPurchase(purchase models.Purchase) (int, error)
	GetPurchase(purchaseId int) (*models.Purchase, error)
	GetPurchases() ([]models.Purchase, error)
	UpdatePurchases(purchaseId int, purchase models.Purchase) error
	DeletePurchase(purchaseId int) error
}

type Stock interface {
	AddStock(stock models.Stock) (int, error)
}

type Service struct {
	Authorisation
	Purchase
	Stock
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Authorisation: NewAuthService(r.Authorisation),
		Purchase:      NewPurchaseService(r.Purchase),
		Stock:         NewStockService(r.Stock),
	}
}
