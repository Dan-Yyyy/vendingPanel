package models

import "time"

type Purchase struct {
	Id           int       `json:"-" db:"id"`
	UserId       int       `json:"user_id" binding:"required"`
	ConsumableId int       `json:"consumable_id" binding:"required"`
	Amount       int       `json:"amount" binding:"required"`
	Price        int       `json:"price" binding:"required"`
	CreatedAt    time.Time `json:"created_at" binding:"required"`
	UpdatedAt    time.Time `json:"updated_at" binding:"required"`
}
