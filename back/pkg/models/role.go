package models

type Role struct {
	Id   int    `json:"-" db:"id"`
	Name string `json:"name" binding:"required,max=255"`
}
