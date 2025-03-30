package mysql

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lecterkn/goat_backend/internal/app/entity"
	"github.com/lecterkn/goat_backend/internal/app/port"
	"github.com/lecterkn/goat_backend/internal/app/repository/mysql/model"
)

type StoryRepositoryImpl struct {
	database *sqlx.DB
}

func NewStoryRepositoryImpl(database *sqlx.DB) port.StoryRepository {
	return &StoryRepositoryImpl{
		database,
	}
}

func (r *StoryRepositoryImpl) Create(ctx context.Context, storyEntity *entity.StoryEntity) error {
	query := `
        INSERT INTO stories(id, category_id, title, episode, description, image_url, created_at, updated_at)
        VALUES(:id, :categoryId, :episode, :description, :imageUrl, :createdAt, :updatedAt)
    `
	return RunInTx(ctx, r.database, func(tx *sqlx.Tx) error {
		queryMap := map[string]any{
			"id":          storyEntity.Id[:],
			"categoryId":  storyEntity.Category.Id[:],
			"episode":     storyEntity.Episode,
			"description": storyEntity.Description,
			"createdAt":   storyEntity.CreatedAt,
			"updatedAt":   storyEntity.UpdatedAt,
		}
		if storyEntity.ImageUrl == nil {
			queryMap["imageUrl"] = sql.NullString{
				String: "",
				Valid:  false,
			}
		} else {
			queryMap["imageUrl"] = *storyEntity.ImageUrl
		}
		_, err := tx.NamedExec(query, queryMap)
		return err
	})
}

func (r *StoryRepositoryImpl) FindAll(ctx context.Context, keyword *string) ([]entity.StoryEntity, error) {
	query := `
        SELECT stories.id, stories.category_id, stories.title, stories.episode, stories.description, stories.image_url, stories.created_at, stories.updated_at,
            categories.name AS category_name, categories.description AS category_desc,  categories.image_url AS category_image_url,
            categories.created_at AS category_created_at, categories.updated_at AS category_updated_at
        FROM stories
        JOIN categories
        ON stories.category_id = categories.id
    `
	if keyword != nil {
		query += `
            WHERE title LIKE ?
        `
	}
	query += "LIMIT 100"
	storyModels := []model.StoryAndCategoryModel{}
	err := RunInTx(ctx, r.database, func(tx *sqlx.Tx) error {
		if keyword != nil {
			return tx.Select(&storyModels, query, *keyword)
		} else {
			return tx.Select(&storyModels, query)
		}
	})
	if err != nil {
		return nil, err
	}
	return r.toEntities(storyModels)
}

func (r *StoryRepositoryImpl) FindById(ctx context.Context, id uuid.UUID) (*entity.StoryEntity, error) {
	query := `
        SELECT stories.id, stories.category_id, stories.title, stories.episode, stories.description, stories.image_url, stories.created_at, stories.updated_at
            categories.name AS category_name, categories.description AS category_desc,  categories.image_url AS category_image_url
            categories.created_at category_created_at, categories.updated_at AS category_updated_at
        FROM stories
        JOIN categories
            ON stories.category_id = categories.id
        WHERE stories.id = ?
        LIMIT 1
    `
	storyModel := model.StoryAndCategoryModel{}
	err := RunInTx(ctx, r.database, func(tx *sqlx.Tx) error {
		return tx.Get(&storyModel, query, id[:])
	})
	if err != nil {
		return nil, err
	}
	return r.toEntity(&storyModel)
}

func (r *StoryRepositoryImpl) toEntities(storyModels []model.StoryAndCategoryModel) ([]entity.StoryEntity, error) {
	storyEntities := []entity.StoryEntity{}
	for _, storyModel := range storyModels {
		storyEntity, err := r.toEntity(&storyModel)
		if err != nil {
			return nil, err
		}
		storyEntities = append(storyEntities, *storyEntity)
	}
	return storyEntities, nil
}

func (r *StoryRepositoryImpl) toEntity(storyModel *model.StoryAndCategoryModel) (*entity.StoryEntity, error) {
	id, err := uuid.FromBytes(storyModel.Id)
	if err != nil {
		return nil, err
	}
	categoryId, err := uuid.FromBytes(storyModel.CategoryId)
	if err != nil {
		return nil, err
	}
	var imageUrl *string = nil
	if storyModel.ImageUrl.Valid {
		imageUrl = &storyModel.ImageUrl.String
	}
	var categoryImageUrl *string = nil
	if storyModel.CategoryImageUrl.Valid {
		categoryImageUrl = &storyModel.CategoryImageUrl.String
	}
	return &entity.StoryEntity{
		Id: id,
		Category: entity.CategoryEntity{
			Id:          categoryId,
			Name:        storyModel.CategoryName,
			Description: storyModel.CategoryDescription,
			ImageUrl:    categoryImageUrl,
			CreatedAt:   storyModel.CategoryCreatedAt,
			UpdatedAt:   storyModel.CategoryUpdatedAt,
		},
		Title:       storyModel.Title,
		Episode:     storyModel.Episode,
		Description: storyModel.Description,
		ImageUrl:    imageUrl,
		CreatedAt:   storyModel.CreatedAt,
		UpdatedAt:   storyModel.UpdatedAt,
	}, nil
}
