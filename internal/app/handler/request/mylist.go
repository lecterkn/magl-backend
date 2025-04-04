package request

import "github.com/google/uuid"

type MyListAddRequest struct {
	StoryId uuid.UUID `json:"storyId" validate:"required"`
	Score   int       `json:"score" validate:"required"`
}

type MyListUpdateRequest struct {
	StoryId uuid.UUID `json:"storyId" validate:"required"`
	Score   int       `json:"score" validate:"required"`
}
