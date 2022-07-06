package models

import "time"

type Purchase struct {
	Id           int       `json:"-" db:"id"`
	UserId       int       `json:"user_id" db:"user_id" binding:"required"`
	ConsumableId int       `json:"consumable_id" db:"consumable_id" binding:"required"`
	Amount       int       `json:"amount" db:"amount" binding:"required"`
	Price        int       `json:"price" db:"price" binding:"required"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
