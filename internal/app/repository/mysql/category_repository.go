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

type CategoryRepositoryImpl struct {
	database *sqlx.DB
}

func NewCategoryRepositoryImpl(database *sqlx.DB) port.CategoryRepository {
	return &CategoryRepositoryImpl{
		database,
	}
}

func (r *CategoryRepositoryImpl) Create(ctx context.Context, categoryEntity *entity.CategoryEntity) error {
	query := `
        INSERT INTO categories(id, name, description, image_url, created_at, updated_at)
        VALUES (:id, :name, :description, :imageUrl, :createdAt, :updatedAt)
    `
	return RunInTx(ctx, r.database, func(tx *sqlx.Tx) error {
		queryMap := map[string]any{
			"id":          categoryEntity.Id[:],
			"name":        categoryEntity.Name,
			"description": categoryEntity.Description,
			"createdAt":   categoryEntity.CreatedAt,
			"updatedAt":   categoryEntity.UpdatedAt,
		}
		if categoryEntity.ImageUrl == nil {
			queryMap["imageUrl"] = sql.NullString{
				String: "",
				Valid:  false,
			}
		} else {
			queryMap["imageUrl"] = *categoryEntity.ImageUrl
		}
		_, err := tx.NamedExec(query, queryMap)
		return err
	})
}

func (r *CategoryRepositoryImpl) FindAll(ctx context.Context, keyword *string) ([]entity.CategoryEntity, error) {
	query := `
        SELECT id, name, description, image_url, created_at, updated_at
        FROM categories
    `
	if keyword != nil {
		query += `
            WHERE name LIKE ?
        `
	}
	query += "LIMIT 100"
	categoryModels := []model.CategoryModel{}
	err := RunInTx(ctx, r.database, func(tx *sqlx.Tx) error {
		if keyword == nil {
			return tx.Select(&categoryModels, query)
		} else {
			return tx.Select(&categoryModels, query, *keyword)
		}
	})
	if err != nil {
		return nil, err
	}
	return r.toEntities(categoryModels)
}

func (r *CategoryRepositoryImpl) FindById(ctx context.Context, id uuid.UUID) (*entity.CategoryEntity, error) {
	query := `
        SELECT id, name, description, image_url, created_at, updated_at
        FROM categories
        WHERE id = ?
        LIMIT 1
    `
	categoryModels := model.CategoryModel{}
	err := RunInTx(ctx, r.database, func(tx *sqlx.Tx) error {
		return tx.Get(&categoryModels, query)
	})
	if err != nil {
		return nil, err
	}
	return r.toEntity(categoryModels)
}

func (r *CategoryRepositoryImpl) toEntities(models []model.CategoryModel) ([]entity.CategoryEntity, error) {
	categoryEntities := []entity.CategoryEntity{}
	for _, categoryModel := range models {
		categoryEntity, err := r.toEntity(*&categoryModel)
		if err != nil {
			return nil, err
		}
		categoryEntities = append(categoryEntities, *categoryEntity)
	}
	return categoryEntities, nil
}

func (r *CategoryRepositoryImpl) toEntity(categoryModel model.CategoryModel) (*entity.CategoryEntity, error) {
	id, err := uuid.FromBytes(categoryModel.Id)
	if err != nil {
		return nil, err
	}
	var imageUrl *string = nil
	if categoryModel.ImageUrl.Valid {
		imageUrl = &categoryModel.ImageUrl.String
	}
	return &entity.CategoryEntity{
		Id:          id,
		Name:        categoryModel.Name,
		Description: categoryModel.Description,
		ImageUrl:    imageUrl,
		CreatedAt:   categoryModel.CreatedAt,
		UpdatedAt:   categoryModel.UpdatedAt,
	}, nil
}
