package model

import (
	"database/sql"
	"time"
)

type StoryModel struct {
	Id          []byte         `db:"id"`
	CategoryId  []byte         `db:"category_id"`
	Title       string         `db:"title"`
	Episode     string         `db:"episode"`
	Description string         `db:"description"`
	ImageUrl    sql.NullString `db:"description"`
	CreatedAt   time.Time      `db:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at"`
}

type StoryAndCategoryModel struct {
	Id                  []byte         `db:"id"`
	Title               string         `db:"title"`
	Episode             string         `db:"episode"`
	Description         string         `db:"description"`
	ImageUrl            sql.NullString `db:"description"`
	CreatedAt           time.Time      `db:"created_at"`
	UpdatedAt           time.Time      `db:"updated_at"`
	CategoryId          []byte         `db:"category_id"`
	CategoryName        string         `db:"category_name"`
	CategoryDescription string         `db:"category_desc"`
	CategoryImageUrl    sql.NullString `db:"category_image_url"`
	CategoryCreatedAt   time.Time      `db:"category_created_at"`
	CategoryUpdatedAt   time.Time      `db:"category_updated_at"`
}
