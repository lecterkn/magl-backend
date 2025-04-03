package input

import "github.com/google/uuid"

type MyListAddInput struct {
	StoryId uuid.UUID
	Score   int
}
