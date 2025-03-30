package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/lecterkn/goat_backend/internal/app/port"
	"github.com/lecterkn/goat_backend/internal/app/usecase/output"
)

type UserUsecase struct {
	userRepository port.UserRepository
}

func NewUserUsecase(
	userRepository port.UserRepository,
) *UserUsecase {
	return &UserUsecase{
		userRepository,
	}
}

// ユーザーをIDから取得
func (u *UserUsecase) GetUser(id uuid.UUID) (*output.UserOutput, error) {
	userEntity, err := u.userRepository.FindById(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return &output.UserOutput{
		Id:        userEntity.Id,
		Name:      userEntity.Name,
		Email:     userEntity.Email,
		Role:      userEntity.Role.GetPermission(),
		CreatedAt: userEntity.CreatedAt,
		UpdatedAt: userEntity.UpdatedAt,
	}, nil
}
