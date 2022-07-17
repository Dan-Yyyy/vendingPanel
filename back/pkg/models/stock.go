package models

import "time"

type Stock struct {
	Id         int       `json:"-" db:"id"`
	PurchaseId int       `json:"purchase_id" binding:"required"`
	CreatedAt  time.Time `json:"created_at"  binding:"required"`
	UpdatedAt  time.Time `json:"updated_at"  binding:"required"`
}

type Consumable struct {
	Id   int    `json:"-" db:"id"`
	Name string `json:"name" binding:"required,max=255"`
}
