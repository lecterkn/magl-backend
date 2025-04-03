//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/lecterkn/goat_backend/internal/app/database"
	"github.com/lecterkn/goat_backend/internal/app/handler"
	"github.com/lecterkn/goat_backend/internal/app/provider"
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
	mysqlRepo.NewStoryRepositoryImpl,
	redisRepo.NewTokenRepositoryImpl,
	mysqlRepo.NewMyListRepositoryImpl,
)

var providerSet = wire.NewSet(
	provider.NewTransactionProviderImpl,
)

var usecaseSet = wire.NewSet(
	usecase.NewAuthorizationUsecase,
	usecase.NewCategoryUsecase,
	usecase.NewStoryUsecase,
	usecase.NewUserUsecase,
	usecase.NewMyListUsecase,
)

var handlerSet = wire.NewSet(
	handler.NewAuthorizationHandler,
	handler.NewCategoryHandler,
	handler.NewStoryHandler,
	handler.NewUserHandler,
	handler.NewMyListHandler,
)

type HandlerSet struct {
	AuthorizationHandler *handler.AuthorizationHandler
	CategoryHandler      *handler.CategoryHandler
	StoryHandler         *handler.StoryHandler
	UserHandler          *handler.UserHandler
	MyListHandler        *handler.MyListHandler
}

func InitializeHandlerSet() *HandlerSet {
	wire.Build(
		databaseSet,
		repositorySet,
		providerSet,
		usecaseSet,
		handlerSet,
		wire.Struct(new(HandlerSet), "*"),
	)
	return nil
}
