package output

import (
	"github.com/google/uuid"
)

type MyListOutput struct {
	UserId  uuid.UUID
	Stories []MyListItemOutput
}

type MyListItemOutput struct {
	StoryId      uuid.UUID
	CategoryId   uuid.UUID
	CategoryName string
	Title        string
	Episode      string
	Description  string
	ImageUrl     *string
	Score        int
}
