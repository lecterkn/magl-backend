package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lecterkn/goat_backend/internal/app/handler/response"
	"github.com/lecterkn/goat_backend/internal/app/usecase"
	"github.com/lecterkn/goat_backend/internal/app/usecase/input"
)

type StoryHandler struct {
	storyUsecase *usecase.StoryUsecase
}

func NewStoryHandler(
	storyUsecase *usecase.StoryUsecase,
) *StoryHandler {
	return &StoryHandler{
		storyUsecase,
	}
}

// @summary		CreateStory
// @description	ストーリーを新規作成
// @tags			story
// @produce		json
// @security		BearerAuth
// @param			image		formData	file	false	"画像ファイル"
// @param			categoryId	formData	string	true	"カテゴリID"
// @param			title		formData	string	true	"ストーリータイトル"
// @param			episode		formData	string	true	"ストーリー区分"
// @param			description	formData	string	true	"ストーリー概要"
// @success		204
// @router			/stories [post]
func (h *StoryHandler) Create(ctx echo.Context) error {
	userId, err := uuid.Parse(ctx.Get("userId").(string))
	if err != nil {
		return err
	}
	categoryId, err := uuid.Parse(ctx.FormValue("categoryId"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "invalid categoryId",
		})
	}
	title := ctx.FormValue("title")
	episode := ctx.FormValue("episode")
	description := ctx.FormValue("description")
	imageFile, _ := ctx.FormFile("image")
	err = h.storyUsecase.CreateStory(userId, input.StoryCreateInput{
		CategoryId:  categoryId,
		Title:       title,
		Episode:     episode,
		Description: description,
		ImageFile:   imageFile,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
	}
	return ctx.NoContent(http.StatusNoContent)
}

// @summary		GetStories
// @description	カテゴリを一覧取得
// @tags			story
// @produce		json
// @param			keyword	query		string	false	"検索キーワード"
// @success		200		{object}	response.StoryListResponse
// @router			/stories [get]
func (h *StoryHandler) GetStories(ctx echo.Context) error {
	var keyword *string = nil
	if ctx.QueryParams().Has("keyword") {
		word := ctx.QueryParam("keyword")
		keyword = &word
	}
	listOutput, err := h.storyUsecase.GetStories(input.StoryQueryInput{
		Keyword: keyword,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	list := []response.StoryResponse{}
	for _, storyOutput := range listOutput {
		list = append(list, response.StoryResponse(storyOutput))
	}
	return ctx.JSON(http.StatusOK, response.StoryListResponse{
		List: list,
	})
}
