package input

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type StoryCreateInput struct {
	CategoryId  uuid.UUID
	Title       string
	Episode     string
	Description string
	ImageFile   *multipart.FileHeader
}

type StoryQueryInput struct {
	Keyword *string
}
