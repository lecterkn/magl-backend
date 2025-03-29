package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CategoryEntity struct {
	Id          uuid.UUID
	Name        string `validate:"required,min=2,max=32"`
	Description string `validate:"required,min=0,max=1023"`
	ImageUrl    *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewCategoryEntity(name, description string, imageUrl *string) (*CategoryEntity, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	categoryEntity := CategoryEntity{
		Id:          id,
		Name:        name,
		Description: description,
		ImageUrl:    imageUrl,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	// バリデーション
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(categoryEntity); err != nil {
		return nil, err
	}
	return &categoryEntity, nil
}
