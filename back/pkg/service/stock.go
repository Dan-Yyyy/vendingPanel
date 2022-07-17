package service

import (
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/models"
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/repository"
)

type StockService struct {
	r repository.Stock
}

func NewStockService(r repository.Stock) *StockService {
	return &StockService{r: r}
}

func (s StockService) AddStock(stock models.Stock) (int, error) {
	return s.r.AddStock(stock)
}
