package models

type User struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name" binding:"required,max=255"`
	Email    string `json:"email" db:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,max=255"`
	RoleId   int    `json:"role_id" db:"role_id" binding:"required"`
}
