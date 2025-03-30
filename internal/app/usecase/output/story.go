package output

import (
	"time"

	"github.com/google/uuid"
)

type StoryOutput struct {
	Id           uuid.UUID
	CategoryId   uuid.UUID
	CategoryName string
	Title        string
	Episode      string
	Description  string
	ImageUrl     *string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
