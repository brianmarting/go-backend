package model

import "github.com/google/uuid"

type User struct {
	Id   int       `db:"id"`
	Uuid uuid.UUID `db:"uuid"`
	Name string    `db:"name"`
}
