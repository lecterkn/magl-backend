package output

import (
	"time"

	"github.com/google/uuid"
)

type UserCreateOutput struct {
	Id        uuid.UUID
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserLoginOutput struct {
	AccessToken  string
	RefreshToken string
}
