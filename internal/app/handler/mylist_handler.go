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
