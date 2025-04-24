package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/lecterkn/goat_backend/internal/app/entity"
	"github.com/lecterkn/goat_backend/internal/app/port"
	"github.com/lecterkn/goat_backend/internal/app/usecase/output"
)

type UserUsecase struct {
	userRepository port.UserRepository
	txProvider     port.TransactionProvider
}

func NewUserUsecase(
	userRepository port.UserRepository,
	txProvider port.TransactionProvider,
) *UserUsecase {
	return &UserUsecase{
		userRepository,
		txProvider,
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

// ユーザーを一覧取得
func (u *UserUsecase) GetUsers(queryUserId uuid.UUID) ([]output.UserOutput, error) {
	userEntities := []entity.UserEntity{}
	err := u.txProvider.Transact(func(ctx context.Context) error {
		// クエリ発行ユーザー取得
		userEntity, err := u.userRepository.FindById(ctx, queryUserId)
		if err != nil {
			return err
		}
		// 権限確認
		err = canQueryUsers(userEntity)
		if err != nil {
			return err
		}
		// ユーザー一覧取得
		userEntities, err = u.userRepository.FindAll(ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	outputList := []output.UserOutput{}
	for _, userEntity := range userEntities {
		outputList = append(outputList, output.UserOutput{
			Id:        userEntity.Id,
			Name:      userEntity.Name,
			Email:     userEntity.Email,
			Role:      userEntity.Role.GetPermission(),
			CreatedAt: userEntity.CreatedAt,
			UpdatedAt: userEntity.UpdatedAt,
		})
	}
	return outputList, nil
}

// ユーザー一覧取得ができるかの権限確認
func canQueryUsers(userEntity *entity.UserEntity) error {
	if userEntity.Role.IsAdministrator() || userEntity.Role.IsRoot() {
		return nil
	}
	return errors.New("invalid permission")
}
