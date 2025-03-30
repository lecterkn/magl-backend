package response

import (
	"time"

	"github.com/google/uuid"
)

type StoryResponse struct {
	Id           uuid.UUID `json:"id" validate:"required"`
	CategoryId   uuid.UUID `json:"categoryId" validate:"required"`
	CategoryName string    `json:"categoryName" validate:"required"`
	Title        string    `json:"title" validate:"required"`
	Episode      string    `json:"episode" validate:"required"`
	Description  string    `json:"description" validate:"required"`
	ImageUrl     *string   `json:"imageUrl" validate:"required"`
	CreatedAt    time.Time `json:"createdAt" validate:"required"`
	UpdatedAt    time.Time `json:"updatedAt" validate:"required"`
}

type StoryListResponse struct {
	List []StoryResponse `json:"list" validate:"required"`
}
