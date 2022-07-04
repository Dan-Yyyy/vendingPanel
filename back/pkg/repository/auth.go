package repository

import (
	"fmt"
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/models"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, email, password_hash, role_id) values ($1, $2, $3, $4) RETURNING id", UsersTable)
	row := r.db.QueryRow(query, user.Name, user.Email, user.Password, user.RoleId)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
