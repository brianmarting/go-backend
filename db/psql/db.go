package psql

import "github.com/jmoiron/sqlx"

func NewDB() (*sqlx.DB, error) {
	return sqlx.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
}
