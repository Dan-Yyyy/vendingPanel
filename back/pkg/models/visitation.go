package models

import "time"

type Visitation struct {
	Id               int  `json:"-" db:"id"`
	VisitationPlanId int  `json:"visitation_plan_id" binding:"required"`
	PointId          int  `json:"point_id" binding:"required"`
	IsVisited        bool `json:"is_visited" binding:"required"`
}

type VisitationResult struct {
	Id          int       `json:"-" db:"id"`
	CreatedAt   time.Time `json:"created_at" binding:"required"`
	UpdatedAt   time.Time `json:"updated_at" binding:"required"`
	CashAmount  int       `json:"cash_amount" binding:"required"`
	CoinsAmount int       `json:"coins_amount" binding:"required"`
	CupsAmount  int       `json:"cups_amount" binding:"required"`
	Comment     string    `json:"comment"`
}

type VisitationPlan struct {
	Id             int       `json:"-" db:"id"`
	VisitationDate time.Time `json:"visitation_date" binding:"required"`
	ResponsibleId  int       `json:"responsible_id" binding:"required"`
}
