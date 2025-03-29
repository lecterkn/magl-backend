package port

import (
	"context"

	"github.com/lecterkn/goat_backend/internal/app/entity"
)

type StoryRepository interface {
	Create(context.Context, *entity.StoryEntity) error
	FindAll(ctx context.Context, keyword *string) ([]entity.StoryEntity, error)
}
