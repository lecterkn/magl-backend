//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/lecterkn/goat_backend/internal/app/database"
	"github.com/lecterkn/goat_backend/internal/app/handler"
	mysqlRepo "github.com/lecterkn/goat_backend/internal/app/repository/mysql"
	redisRepo "github.com/lecterkn/goat_backend/internal/app/repository/redis"
	"github.com/lecterkn/goat_backend/internal/app/usecase"
)

var databaseSet = wire.NewSet(
	// *sqlx.DB
	database.GetMySQLConnection,
	// *redis.Client
	database.GetRedisClient,
)

var repositorySet = wire.NewSet(
	mysqlRepo.NewUserRepositoryImpl,
	mysqlRepo.NewCategoryRepositoryImpl,
	redisRepo.NewTokenRepositoryImpl,
)

var usecaseSet = wire.NewSet(
	usecase.NewAuthorizationUsecase,
	usecase.NewCategoryUsecase,
)

var handlerSet = wire.NewSet(
	handler.NewAuthorizationHandler,
	handler.NewCategoryHandler,
)

type HandlerSet struct {
	AuthorizationHandler *handler.AuthorizationHandler
	CategoryHandler      *handler.CategoryHandler
}

func InitializeHandlerSet() *HandlerSet {
	wire.Build(
		databaseSet,
		repositorySet,
		usecaseSet,
		handlerSet,
		wire.Struct(new(HandlerSet), "*"),
	)
	return nil
}
