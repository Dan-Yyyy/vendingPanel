package repository

import (
	"fmt"
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/models"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func (r *AuthPostgres) GetUserById(userId int) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id, name, email, role_id FROM %s WHERE id=$1", UsersTable)
	err := r.db.Get(&user, query, userId)

	return user, err
}

func NewAuth(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf(
		"INSERT INTO %s (name, email, password_hash, role_id) SELECT $1, $2, $3, $4 WHERE NOT EXISTS (SELECT id FROM %s WHERE email = $5) RETURNING id",
		UsersTable, UsersTable)
	row := r.db.QueryRow(query, user.Name, user.Email, user.Password, user.RoleId, user.Email)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(email string, passwordHash string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password_hash=$2", UsersTable)
	err := r.db.Get(&user, query, email, passwordHash)

	return user, err
}
