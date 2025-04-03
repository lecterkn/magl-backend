package request

import "github.com/google/uuid"

type MyListAddRequest struct {
	StoryId uuid.UUID `json:"storyId"`
	Score   int       `json:"score"`
}
