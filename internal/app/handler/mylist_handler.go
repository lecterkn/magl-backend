package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lecterkn/goat_backend/internal/app/handler/request"
	"github.com/lecterkn/goat_backend/internal/app/handler/response"
	"github.com/lecterkn/goat_backend/internal/app/usecase"
	"github.com/lecterkn/goat_backend/internal/app/usecase/input"
)

type MyListHandler struct {
	mylistUsecase *usecase.MyListUsecase
}

func NewMyListHandler(mylistUsecase *usecase.MyListUsecase) *MyListHandler {
	return &MyListHandler{
		mylistUsecase,
	}
}

//	@summary		AddMyList
//	@description	マイリストにストーリーを追加
//	@tags			mylist
//	@produce		json
//	@security		BearerAuth
//	@param			request	body	request.MyListAddRequest	true	"マイリスト追加リクエスト"
//	@success		204
//	@router			/mylists [post]
func (h *MyListHandler) AddMyList(ctx echo.Context) error {
	userId, err := uuid.Parse(ctx.Get("userId").(string))
	if err != nil {
		return err
	}
	addRequest := request.MyListAddRequest{}
	if err := ctx.Bind(&addRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "invalid request",
		})
	}
	err = h.mylistUsecase.AddStoryToMyList(
		userId, input.MyListAddInput{
			StoryId: addRequest.StoryId,
			Score:   addRequest.Score,
		},
	)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
	}
	return ctx.NoContent(http.StatusNoContent)
}

//	@summary		UpdateMyList
//	@description	マイリストのストーリーを更新
//	@tags			mylist
//	@produce		json
//	@security		BearerAuth
//	@param			request	body	request.MyListUpdateRequest	true	"マイリスト更新リクエスト"
//	@success		204
//	@router			/mylists [patch]
func (h *MyListHandler) UpdateMyList(ctx echo.Context) error {
	userId, err := uuid.Parse(ctx.Get("userId").(string))
	if err != nil {
		return err
	}
	updateRequest := request.MyListUpdateRequest{}
	if err := ctx.Bind(&updateRequest); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: "invalid request body",
		})
	}
	err = h.mylistUsecase.UpdateScore(userId, input.MyListUpdateInput{
		StoryId: updateRequest.StoryId,
		Score:   updateRequest.Score,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
	}
	return ctx.NoContent(http.StatusNoContent)
}

//	@summary		GetMyList
//	@description	マイリストを取得
//	@tags			mylist
//	@produce		json
//	@security		BearerAuth
//	@success		200	{object}	response.MyListListResponse
//	@router			/mylists [get]
func (h *MyListHandler) GetMyList(ctx echo.Context) error {
	userId, err := uuid.Parse(ctx.Get("userId").(string))
	if err != nil {
		return err
	}
	mylistOutput, err := h.mylistUsecase.GetMylist(userId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
	}
	list := []response.MyListResponse{}
	for _, story := range mylistOutput.Stories {
		list = append(list, response.MyListResponse{
			Id:           story.StoryId,
			Title:        story.Title,
			Episode:      story.Episode,
			Description:  story.Description,
			ImageUrl:     story.ImageUrl,
			CategoryId:   story.CategoryId,
			CategoryName: story.CategoryName,
			Score:        story.Score,
		})
	}
	return ctx.JSON(http.StatusOK, response.MyListListResponse{
		List: list,
	})
}

//	@summary		RemoveFromMyList
//	@description	マイリストからストーリーを削除
//	@tags			mylist
//	@produce		json
//	@security		BearerAuth
//	@param			storyId	path	string	true	"ストーリーID"
//	@success		204
//	@router			/mylists/{storyId} [delete]
func (h *MyListHandler) RemoveFromMyList(ctx echo.Context) error {
	userId, err := uuid.Parse(ctx.Get("userId").(string))
	if err != nil {
		return err
	}
	storyId, err := uuid.Parse(ctx.Param("storyId"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "invalid storyId",
		})
	}
	err = h.mylistUsecase.RemoveFromList(userId, storyId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
	}
	return ctx.NoContent(http.StatusNoContent)
}
