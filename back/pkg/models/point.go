package models

type Point struct {
	Id      int    `json:"-" db:"id"`
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
}

type CoffeeMachineType struct {
	Id   int    `json:"-" db:"id"`
	Name string `json:"name" binding:"required"`
}

type CoffeeMachinePoint struct {
	CoffeeMachineTypeId int `json:"coffee_machine_type_id" binding:"required"`
	PointId             int `json:"point_id" binding:"required"`
}
