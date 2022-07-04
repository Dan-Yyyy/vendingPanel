package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable                  = "users"
	additionalExpensesTable     = "additional_expenses"
	additionalExpenseTypesTable = "additional_expense_types"
	pointsTable                 = "points"
	coffeeMachineTypesTable     = "coffee_machine_types"
	coffeeMachinePointsTable    = "coffee_machine_points"
	purchasesTable              = "purchases"
	reportsTable                = "reports"
	reportTypesTable            = "report_types"
	reportStocksTable           = "report_stocks"
	rolesTable                  = "roles"
	stocksTable                 = "stocks"
	consumablesTable            = "consumables"
	visitationsTable            = "visitations"
	visitationResultsTable      = "visitation_results"
	visitationPlansTable        = "visitation_plans"
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
