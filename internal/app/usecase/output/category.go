package output

import (
	"time"

	"github.com/google/uuid"
	"github.com/lecterkn/goat_backend/internal/app/entity"
)

type CategoryQueryOutput struct {
	Id          uuid.UUID
	Name        string
	Description string
	ImageUrl    *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func ToOutput(category entity.CategoryEntity) CategoryQueryOutput {
	imageUrl := &category.ImageUrl.Url
	if category.ImageUrl.Availability == entity.Unavailable {
		imageUrl = nil
	}
	return CategoryQueryOutput{
		Id:          category.Id,
		Name:        category.Name,
		Description: category.Description,
		ImageUrl:    imageUrl,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}
}
