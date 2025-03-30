package model

import "time"

type UserModel struct {
	Id        []byte    `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  []byte    `db:"password"`
	Role      int       `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
