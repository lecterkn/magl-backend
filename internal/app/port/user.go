package port

import (
	"context"

	"github.com/lecterkn/goat_backend/internal/app/entity"
)

type UserRepository interface {
	Create(context.Context, *entity.UserEntity) error
	FindByName(context.Context, string) (*entity.UserEntity, error)
}
