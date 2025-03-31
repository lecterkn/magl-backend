package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type StoryEntity struct {
	Id          uuid.UUID
	Category    CategoryEntity
	Title       string `validate:"required,min=3,max=64"`
	Episode     string `validate:"required,min=1,max=64"`
	Description string `validate:"min=0,max=1024"`
	ImageUrl    *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewStoryEntity(categoryEntity CategoryEntity, title, episode, description string, imageUrl *string) (*StoryEntity, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	storyEntity := StoryEntity{
		Id:          id,
		Category:    categoryEntity,
		Title:       title,
		Episode:     episode,
		Description: description,
		ImageUrl:    imageUrl,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(storyEntity); err != nil {
		return nil, err
	}
	return &storyEntity, nil
}
