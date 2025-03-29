package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lecterkn/goat_backend/internal/app/handler/response"
	"github.com/lecterkn/goat_backend/internal/app/usecase"
	"github.com/lecterkn/goat_backend/internal/app/usecase/input"
)

type CategoryHandler struct {
	categoryUsecase *usecase.CategoryUsecase
}

func NewCategoryHandler(
	categoryUsecase *usecase.CategoryUsecase,
) *CategoryHandler {
	return &CategoryHandler{
		categoryUsecase,
	}
}

//	@summary		CreateCategory
//	@description	カテゴリを新規作成
//	@tags			category
//	@produce		json
//	@security		BearerAuth
//	@param			image		formData	file	true	"画像ファイル"
//	@param			name		formData	string	true	"カテゴリ名"
//	@param			description	formData	string	true	"カテゴリ概要"
//	@success		204
//	@router			/categories [post]
func (h *CategoryHandler) Create(ctx echo.Context) error {
	name := ctx.FormValue("name")
	description := ctx.FormValue("description")
	imageFile, err := ctx.FormFile("image")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "invalid image",
		})
	}
	err = h.categoryUsecase.CreateCategory(input.CategoryCreateInput{
		Name:        name,
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

//	@summary		GetCategories
//	@description	カテゴリを一覧取得
//	@tags			category
//	@produce		json
//	@param			keyword	query		string	false	"検索キーワード"
//	@success		200		{object}	response.CategoryListResponse
//	@router			/categories [get]
func (h *CategoryHandler) FindAll(ctx echo.Context) error {
	var keyword *string = nil
	if ctx.QueryParams().Has("keyword") {
		word := ctx.QueryParam("keyword")
		keyword = &word
	}
	outputList, err := h.categoryUsecase.GetCategories(input.CategoryQueryInput{
		Keyword: keyword,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
	}
	list := []response.CategoryResponse{}
	for _, output := range outputList {
		list = append(list, response.CategoryResponse(output))
	}
	return ctx.JSON(http.StatusOK, response.CategoryListResponse{
		List: list,
	})
}
