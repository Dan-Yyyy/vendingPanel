package models

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required,max=255"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,max=255"`
	RoleId   int    `json:"role_id" binding:"required"`
}
