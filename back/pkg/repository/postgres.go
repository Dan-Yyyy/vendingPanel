package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	UsersTable                  = "users"
	AdditionalExpensesTable     = "additional_expenses"
	AdditionalExpenseTypesTable = "additional_expense_types"
	PointsTable                 = "points"
	CoffeeMachineTypesTable     = "coffee_machine_types"
	CoffeeMachinePointsTable    = "coffee_machine_points"
	PurchasesTable              = "purchases"
	ReportsTable                = "reports"
	ReportTypesTable            = "report_types"
	ReportStocksTable           = "report_stocks"
	RolesTable                  = "roles"
	StocksTable                 = "stocks"
	ConsumablesTable            = "consumables"
	VisitationsTable            = "visitations"
	VisitationResultsTable      = "visitation_results"
	VisitationPlansTable        = "visitation_plans"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
