package mysql

import (
	"context"
	"fmt"

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
        INSERT INTO users(id, name, email, password, created_at, updated_at)
        VALUES (:id, :name, :email, :password, :createdAt, :updatedAt)
    `
	return RunInTx(ctx, r.database, func(tx *sqlx.Tx) error {
		_, err := tx.NamedExec(query, map[string]any{
			"id":        userEntity.Id[:],
			"name":      userEntity.Name,
			"email":     userEntity.Email,
			"password":  userEntity.Password,
			"createdAt": userEntity.CreatedAt,
			"updatedAt": userEntity.UpdatedAt,
		})
		return err
	})
}

func (r *UserRepositoryImpl) FindByName(ctx context.Context, name string) (*entity.UserEntity, error) {
	query := `
        SELECT id, name, email, password, created_at, updated_at
        FROM users
        WHERE name = ?
    `
	userModel := model.UserModel{}
	err := RunInTx(ctx, r.database, func(tx *sqlx.Tx) error {
		fmt.Println(name)
		return tx.Get(&userModel, query, name)
	})
	if err != nil {
		return nil, err
	}
	return r.toEntity(&userModel)
}

func (r *UserRepositoryImpl) toEntity(userModel *model.UserModel) (*entity.UserEntity, error) {
	id, err := uuid.FromBytes(userModel.Id)
	if err != nil {
		return nil, err
	}
	return &entity.UserEntity{
		Id:        id,
		Name:      userModel.Name,
		Email:     userModel.Email,
		Password:  userModel.Password,
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,
	}, nil
}
