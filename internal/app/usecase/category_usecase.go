package usecase

import (
	"context"

	"github.com/lecterkn/goat_backend/internal/app/entity"
	"github.com/lecterkn/goat_backend/internal/app/port"
	"github.com/lecterkn/goat_backend/internal/app/usecase/input"
	"github.com/lecterkn/goat_backend/internal/app/usecase/output"
)

type CategoryUsecase struct {
	categoryRepository port.CategoryRepository
}

func NewCategoryUsecase(
	categoryRepository port.CategoryRepository,
) *CategoryUsecase {
	return &CategoryUsecase{
		categoryRepository,
	}
}

// カテゴリを新規作成
func (u *CategoryUsecase) CreateCategory(cmd input.CategoryCreateInput) error {
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
