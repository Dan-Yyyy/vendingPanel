package models

type AdditionalExpense struct {
	Id                      int    `json:"-" db:"id"`
	AdditionalExpenseTypeId int    `json:"additional_expense_type_id" binding:"required"`
	Name                    string `json:"name" binding:"required,max=255"`
	Amount                  int    `json:"amount" binding:"required"`
}

type AdditionalExpenseType struct {
	Id     int `json:"-" db:"id"`
	Amount int `json:"amount" binding:"required"`
}
