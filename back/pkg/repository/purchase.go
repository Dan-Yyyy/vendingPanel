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

func (r PurchasePostgres) GetPurchase(purchaseId int) (*models.Purchase, error) {
	var purchase models.Purchase
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", PurchasesTable)
	err := r.db.Get(&purchase, query, purchaseId)

	if err != nil {
		return nil, err
	}

	return &purchase, nil
}

func (r PurchasePostgres) GetPurchases() ([]models.Purchase, error) {
	var purchases []models.Purchase

	query := fmt.Sprintf("SELECT * FROM %s", PurchasesTable)
	err := r.db.Select(&purchases, query)
	if err != nil {
		return nil, err
	}

	return purchases, nil
}
