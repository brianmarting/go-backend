package model

import "github.com/google/uuid"

type Crypto struct {
	Id          int       `db:"id"`
	Uuid        uuid.UUID `db:"uuid"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
}
