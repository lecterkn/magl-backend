package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lecterkn/goat_backend/docs"
	"github.com/lecterkn/goat_backend/internal/app/di"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title						MyAnimeGameList
// @version					1.0
// @description				MyAnimeGameList API Server
// @host						localhost:8993
// @BasePath					/api
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	err := godotenv.Load()
	if err != nil {
		panic("\".env\"が見つかりません")
	}
	port, ok := os.LookupEnv("ECHO_SERVER_PORT")
	if !ok {
		panic("環境変数に\"ECHO_SERVER_PORT\"が設定されていません")
	}
	app := echo.New()
	app.Use(middleware.Logger())
	setRouting(app)
	app.Logger.Fatal(app.Start(":" + port))
}

// ルーティングの設定
func setRouting(app *echo.Echo) {
	app.GET("/swagger/*", echoSwagger.WrapHandler)

	handlerSet := di.InitializeHandlerSet()

	api := app.Group("api")
	auth := api.Group("")
	auth.Use(handlerSet.AuthorizationHandler.Authorization)

	// Authorization
	api.POST("/signup", handlerSet.AuthorizationHandler.SignUp)
	api.POST("/signin", handlerSet.AuthorizationHandler.SignIn)

	// Category
	auth.POST("/categories", handlerSet.CategoryHandler.Create)
	api.GET("/categories", handlerSet.CategoryHandler.GetCategories)

	// Story
	auth.POST("/stories", handlerSet.StoryHandler.Create)
	api.GET("/stories", handlerSet.StoryHandler.GetStories)
}
