package models

import "time"

type Report struct {
	Id                 int       `json:"-" db:"id"`
	UserId             int       `json:"role_id" binding:"required"`
	ReportTypeId       int       `json:"report_type_id" binding:"required"`
	VisitationId       int       `json:"visitation_id" binding:"required"`
	CreatedAt          time.Time `json:"created_at" binding:"required"`
	UpdatedAt          time.Time `json:"updated_at" binding:"required"`
	IsProductAdded     bool      `json:"is_product_added" binding:"required"`
	IsEncasement       bool      `json:"is_encasement" binding:"required"`
	Password           string    `json:"password" binding:"required"`
	VisitationResultId int       `json:"visitation_result_id" binding:"required"`
}

type ReportType struct {
	Id   int    `json:"-" db:"id"`
	Name string `json:"name" binding:"required"`
}

type ReportStock struct {
	ReportId int `json:"report_id" binding:"required"`
	StockId  int `json:"stock_id" binding:"required"`
	Amount   int `json:"amount" binding:"required"`
}
