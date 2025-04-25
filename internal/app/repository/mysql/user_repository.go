package mysql

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lecterkn/goat_backend/internal/app/entity"
	"github.com/lecterkn/goat_backend/internal/app/port"
	"github.com/lecterkn/goat_backend/internal/app/repository/mysql/model"
)

type UserRepositoryImpl struct {
	database *sqlx.DB
}

func NewUserRepositoryImpl(database *sqlx.DB) port.UserRepository {
	return &UserRepositoryImpl{
		database,
	}
}

func (r *UserRepositoryImpl) Create(ctx context.Context, userEntity *entity.UserEntity) error {
	query := `
        INSERT INTO users(id, name, email, password, role, created_at, updated_at)
        VALUES (:id, :name, :email, :password, :role, :createdAt, :updatedAt)
    `
	return RunInTx(ctx, r.database, func(tx *sqlx.Tx) error {
		_, err := tx.NamedExec(query, map[string]any{
			"id":        userEntity.Id[:],
			"name":      userEntity.Name,
			"email":     userEntity.Email,
			"role":      userEntity.Role.Permission,
			"password":  userEntity.Password,
			"createdAt": userEntity.CreatedAt,
			"updatedAt": userEntity.UpdatedAt,
		})
		return err
	})
}

func (r *UserRepositoryImpl) Update(ctx context.Context, userEntity *entity.UserEntity) error {
	query := `
		UPDATE users
		SET name = :name, email = :email, password = :password, role = :role, updated_at = :updatedAt
		WHERE id = :id
    `
	return RunInTx(ctx, r.database, func(tx *sqlx.Tx) error {
		_, err := tx.NamedExec(query, map[string]any{
			"id":        userEntity.Id[:],
			"name":      userEntity.Name,
			"email":     userEntity.Email,
			"role":      userEntity.Role.Permission,
			"password":  userEntity.Password,
			"updatedAt": userEntity.UpdatedAt,
		})
		return err
	})
}

func (r *UserRepositoryImpl) FindByName(ctx context.Context, name string) (*entity.UserEntity, error) {
	query := `
        SELECT id, name, email, password, role, created_at, updated_at
        FROM users
        WHERE name = ?
    `
	userModel := model.UserModel{}
	err := RunInTx(ctx, r.database, func(tx *sqlx.Tx) error {
		return tx.Get(&userModel, query, name)
	})
	if err != nil {
		return nil, err
	}
	return r.toEntity(&userModel)
}

func (r *UserRepositoryImpl) FindById(ctx context.Context, id uuid.UUID) (*entity.UserEntity, error) {
	query := `
        SELECT id, name, email, password, role, created_at, updated_at
        FROM users
        WHERE id = ?
    `
	userModel := model.UserModel{}
	err := RunInTx(ctx, r.database, func(tx *sqlx.Tx) error {
		return tx.Get(&userModel, query, id[:])
	})
	if err != nil {
		return nil, err
	}
	return r.toEntity(&userModel)
}

func (r *UserRepositoryImpl) FindAll(ctx context.Context) ([]entity.UserEntity, error) {
	query := `
        SELECT id, name, email, password, role, created_at, updated_at
        FROM users
    `
	userModels := []model.UserModel{}
	err := RunInTx(ctx, r.database, func(tx *sqlx.Tx) error {
		return tx.Select(&userModels, query)
	})
	if err != nil {
		return nil, err
	}

	return r.toEntities(userModels)
}

func (r *UserRepositoryImpl) toEntities(userModels []model.UserModel) ([]entity.UserEntity, error) {
	userEntities := []entity.UserEntity{}
	for _, userModel := range userModels {
		userEntity, err := r.toEntity(&userModel)
		if err != nil {
			return nil, err
		}
		userEntities = append(userEntities, *userEntity)
	}
	return userEntities, nil
}

func (r *UserRepositoryImpl) toEntity(userModel *model.UserModel) (*entity.UserEntity, error) {
	id, err := uuid.FromBytes(userModel.Id)
	if err != nil {
		return nil, err
	}
	return &entity.UserEntity{
		Id:       id,
		Name:     userModel.Name,
		Email:    userModel.Email,
		Password: userModel.Password,
		Role: &entity.RoleEntity{
			Permission: entity.Permission(userModel.Role),
		},
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,
	}, nil
}
