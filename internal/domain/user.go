package domain

import (
	"database/sql"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID    `db:"id"`
	Role      string       `db:"role"`
	Email     string       `db:"email"`
	Password  string       `db:"password"`
	Name      string       `db:"name"`
	Surname   string       `db:"surname"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	//to add (model) ads: []Ad
	//to add (model)favourites: []Favourites
}
