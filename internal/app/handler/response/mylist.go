package response

import (
	"github.com/google/uuid"
)

type MyListListResponse struct {
	List []MyListResponse `json:"list" validate:"required"`
}

type MyListResponse struct {
	Id           uuid.UUID `json:"id" validate:"required"`
	CategoryId   uuid.UUID `json:"categoryId" validate:"required"`
	CategoryName string    `json:"categoryName" validate:"required"`
	Title        string    `json:"title" validate:"required"`
	Episode      string    `json:"episode" validate:"required"`
	Description  string    `json:"description" validate:"required"`
	ImageUrl     *string   `json:"imageUrl" validate:"required"`
	Score        int       `json:"score" validate:"required"`
}
