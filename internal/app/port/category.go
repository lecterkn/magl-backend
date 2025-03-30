package port

import (
	"context"

	"github.com/google/uuid"
	"github.com/lecterkn/goat_backend/internal/app/entity"
)

type CategoryRepository interface {
	Create(context.Context, *entity.CategoryEntity) error
	FindAll(ctx context.Context, keyword *string) ([]entity.CategoryEntity, error)
	FindById(ctx context.Context, id uuid.UUID) (*entity.CategoryEntity, error)
}
