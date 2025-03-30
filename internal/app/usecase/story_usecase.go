package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/lecterkn/goat_backend/internal/app/entity"
	"github.com/lecterkn/goat_backend/internal/app/port"
	"github.com/lecterkn/goat_backend/internal/app/usecase/input"
)

type StoryUsecase struct {
	userRepository     port.UserRepository
	storyRepository    port.StoryRepository
	categoryRepository port.CategoryRepository
}

func NewStoryUsecase(
	userRepository port.UserRepository,
	storyRepository port.StoryRepository,
	categoryRepository port.CategoryRepository,
) *StoryUsecase {
	return &StoryUsecase{
		userRepository,
		storyRepository,
		categoryRepository,
	}
}

func (u *StoryUsecase) CreateStory(userId uuid.UUID, cmd input.StoryCreateInput) error {
	userEntity, err := u.userRepository.FindById(context.Background(), userId)
	if err != nil {
		return err
	}
	if !u.canCreateStory(userEntity) {
		return errors.New("invalid permission")
	}
	categoryEntity, err := u.categoryRepository.FindById(context.Background(), cmd.CategoryId)
	if err != nil {
		return err
	}
	var imageUrl *string = nil
	// TODO: 画像ファイルをアップロード
	if cmd.ImageFile != nil {
	}
	storyEntity, err := entity.NewStoryEntity(*categoryEntity, cmd.Title, cmd.Episode, cmd.Description, imageUrl)
	if err != nil {
		return err
	}
	return u.storyRepository.Create(context.Background(), storyEntity)
}

func (u *StoryUsecase) canCreateStory(userEntity *entity.UserEntity) bool {
	return userEntity.Role.IsAdministrator() || userEntity.Role.IsRoot()
}
