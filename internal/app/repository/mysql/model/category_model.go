package model

import (
	"database/sql"
	"time"
)

type CategoryModel struct {
	Id          []byte
	Name        string
	Description string
	ImageUrl    sql.NullString
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
