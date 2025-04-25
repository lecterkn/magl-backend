package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/lecterkn/goat_backend/internal/app/entity"
	"github.com/lecterkn/goat_backend/internal/app/port"
	"github.com/lecterkn/goat_backend/internal/app/usecase/input"
	"github.com/lecterkn/goat_backend/internal/app/usecase/output"
)

type CategoryUsecase struct {
	userRepository     port.UserRepository
	categoryRepository port.CategoryRepository
}

func NewCategoryUsecase(
	userRepository port.UserRepository,
	categoryRepository port.CategoryRepository,
) *CategoryUsecase {
	return &CategoryUsecase{
		userRepository,
		categoryRepository,
	}
}

// カテゴリを新規作成
func (u *CategoryUsecase) CreateCategory(userId uuid.UUID, cmd input.CategoryCreateInput) error {
	userEntity, err := u.userRepository.FindById(context.Background(), userId)
	if err != nil {
		return err
	}
	if !u.canCreateCategory(userEntity) {
		return errors.New("permission error")
	}
	var imageUrl *string = nil
	// TODO: ファイルアップロード
	if cmd.ImageFile != nil {
	}
	categoryEntity, err := entity.NewCategoryEntity(
		cmd.Name, cmd.Description, imageUrl,
	)
	if err != nil {
		return err
	}
	return u.categoryRepository.Create(context.Background(), categoryEntity)
}

// カテゴリを一覧取得
func (u *CategoryUsecase) GetCategories(cmd input.CategoryQueryInput) ([]output.CategoryQueryOutput, error) {
	categorieEntities, err := u.categoryRepository.FindAll(context.Background(), cmd.Keyword)
	if err != nil {
		return nil, err
	}
	outputList := []output.CategoryQueryOutput{}
	for _, categoryEntity := range categorieEntities {
		outputList = append(outputList, output.CategoryQueryOutput(categoryEntity))
	}
	return outputList, nil
}

func (u *CategoryUsecase) canCreateCategory(userEntity *entity.UserEntity) bool {
	return userEntity.Role.IsModerator() || userEntity.Role.IsAdministrator() || userEntity.Role.IsRoot()
}
