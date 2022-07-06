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

func NewPurchaseService(r repository.Purchase) *PurchaseService {
	return &PurchaseService{r: r}
}
