package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/lecterkn/goat_backend/internal/app/common"
)

type UserEntity struct {
	Id        uuid.UUID
	Name      string `validate:"required,min=4,max=32"`
	Email     string `validate:"required,email"`
	Password  []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserEntity(name, email, password string) (*UserEntity, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	hashPass, err := common.EncryptPassword(password)
	if err != nil {
		return nil, err
	}
	userEntity := UserEntity{
		Id:        id,
		Name:      name,
		Email:     email,
		Password:  hashPass,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(userEntity); err != nil {
		return nil, err
	}
	return &userEntity, nil
}
