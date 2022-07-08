package service

import (
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/models"
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/repository"
)

type PurchaseService struct {
	r repository.Purchase
}

func (s PurchaseService) AddPurchase(purchase models.Purchase) (int, error) {
	return s.r.AddPurchase(purchase)
}

func (s PurchaseService) GetPurchase(purchaseId int) (*models.Purchase, error) {
	return s.r.GetPurchase(purchaseId)
}

func (s PurchaseService) GetPurchases() ([]models.Purchase, error) {
	return s.r.GetPurchases()
}

func NewPurchaseService(r repository.Purchase) *PurchaseService {
	return &PurchaseService{r: r}
}
