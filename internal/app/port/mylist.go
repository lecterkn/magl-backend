package port

import (
	"context"

	"github.com/google/uuid"
	"github.com/lecterkn/goat_backend/internal/app/entity"
)

type MyListRepository interface {
	Save(context.Context, *entity.MyListEntity) error
	FindByUserId(context.Context, uuid.UUID) (*entity.MyListEntity, error)
}
