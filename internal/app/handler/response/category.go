package response

import (
	"time"

	"github.com/google/uuid"
)

type CategoryResponse struct {
	Id          uuid.UUID `json:"id" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	ImageUrl    *string   `json:"imageUrl"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type CategoryListResponse struct {
	List []CategoryResponse `json:"list" validate:"required"`
}
