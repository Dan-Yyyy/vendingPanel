package models

type Point struct {
	Id      int    `json:"-" db:"id"`
	Name    string `json:"name" binding:"required,max=255"`
	Address string `json:"address" binding:"required,max=255"`
}

type CoffeeMachineType struct {
	Id   int    `json:"-" db:"id"`
	Name string `json:"name" binding:"required,max=255"`
}

type CoffeeMachinePoint struct {
	CoffeeMachineTypeId int `json:"coffee_machine_type_id" binding:"required"`
	PointId             int `json:"point_id" binding:"required"`
}
