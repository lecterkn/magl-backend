package model

import (
	"database/sql"
	"time"
)

type MyListModel struct {
	Score               int            `db:"score"`
	StoryId             []byte         `db:"story_id"`
	StoryTitle          string         `db:"story_title"`
	StoryEpisode        string         `db:"story_episode"`
	StoryDescription    string         `db:"story_desc"`
	StoryImageUrl       sql.NullString `db:"story_image_url"`
	StoryCreatedAt      time.Time      `db:"story_created_at"`
	StoryUpdatedAt      time.Time      `db:"story_updated_at"`
	CategoryId          []byte         `db:"category_id"`
	CategoryName        string         `db:"category_name"`
	CategoryDescription string         `db:"category_desc"`
	CategoryImageUrl    sql.NullString `db:"category_image_url"`
	CategoryCreatedAt   time.Time      `db:"category_created_at"`
	CategoryUpdatedAt   time.Time      `db:"category_updated_at"`
}
