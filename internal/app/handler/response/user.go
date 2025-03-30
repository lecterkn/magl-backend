package response

import (
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	Id        uuid.UUID `json:"id" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"required"`
	Role      string    `json:"role" validare:"required"`
	CreatedAt time.Time `json:"createdAt" validate:"required"`
	UpdatedAt time.Time `json:"updatedAt" validate:"required"`
}

type UserSigninResponse struct {
	AccessToken  string `json:"accessToken" validate:"required"`
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type RefreshResponse struct {
	AccessToken string    `json:"accessToken" validare:"required"`
	ExpiresIn   time.Time `json:"expiresIn" validare:"required"`
}
