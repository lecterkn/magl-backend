package model

import (
	"database/sql"
	"time"
)

type CategoryModel struct {
	Id          []byte         `db:"id"`
	Name        string         `db:"name"`
	Description string         `db:"description"`
	ImageUrl    sql.NullString `db:"image_url"`
	CreatedAt   time.Time      `db:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at"`
}
