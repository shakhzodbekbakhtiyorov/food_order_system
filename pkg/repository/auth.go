package repository

import (
	"fmt"
	"gogogo"

	"github.com/jmoiron/sqlx"
)

type Auth struct {
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) *Auth {
	return &Auth{db: db}
}

func (r *Auth) CreateUser(user gogogo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, login, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Login, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Auth) GetUser(login, password string) (gogogo.User, error) {
	var user gogogo.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE login=$1 and password_hash=$2", usersTable)
	err := r.db.Get(&user, query, login, password)
	return user, err
}
