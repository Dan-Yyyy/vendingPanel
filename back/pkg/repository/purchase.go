package repository

import (
	"fmt"
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/models"
	"github.com/jmoiron/sqlx"
)

type PurchasePostgres struct {
	db *sqlx.DB
}

func NewPurchase(db *sqlx.DB) *PurchasePostgres {
	return &PurchasePostgres{db: db}
}

func (r PurchasePostgres) AddPurchase(purchase models.Purchase) (int, error) {
	var id int
	query := fmt.Sprintf(
		"INSERT INTO %s (user_id, consumable_id, amount, price, created_at, updated_at) VALUES($1,$2,$3,$4,$5,$6) RETURNING id",
		PurchasesTable)
	row := r.db.QueryRow(query, purchase.UserId, purchase.ConsumableId, purchase.Amount, purchase.Price, purchase.CreatedAt, purchase.UpdatedAt)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
