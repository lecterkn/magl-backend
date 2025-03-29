package port

import (
	"github.com/google/uuid"
	"github.com/lecterkn/goat_backend/internal/app/entity"
)

type TokenRepository interface {
	SaveRefreshToken(*entity.RefreshTokenEntity) error
	FindRefreshTokenByUserId(uuid.UUID) ([]entity.RefreshTokenEntity, error)
}
