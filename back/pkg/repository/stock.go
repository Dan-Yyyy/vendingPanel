package repository

import (
	"fmt"
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/models"
	"github.com/jmoiron/sqlx"
)

type StockPostgres struct {
	db *sqlx.DB
}

func NewStock(db *sqlx.DB) *StockPostgres {
	return &StockPostgres{db: db}
}

func (r StockPostgres) AddStock(stock models.Stock) (int, error) {
	var id int
	query := fmt.Sprintf(
		"INSERT INTO %s (purchase_id, created_at, updated_at) VALUES($1,$2,$3) RETURNING id",
		StocksTable)
	row := r.db.QueryRow(query, stock.PurchaseId, stock.CreatedAt, stock.UpdatedAt)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
