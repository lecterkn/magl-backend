package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lecterkn/goat_backend/internal/app/handler/response"
	"github.com/lecterkn/goat_backend/internal/app/usecase"
)

type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewUserHandler(
	userUsecase *usecase.UserUsecase,
) *UserHandler {
	return &UserHandler{
		userUsecase,
	}
}

//	@summary		GetMe
//	@description	自身のユーザー情報を取得する
//	@tags			user
//	@produce		json
//	@success		200	{object}	response.UserResponse
//	@router			/me [get]
func (h *UserHandler) GetMe(ctx echo.Context) error {
	userId, err := uuid.Parse(ctx.Get("userId").(string))
	if err != nil {
		return err
	}
	userOutput, err := h.userUsecase.GetUser(userId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, response.UserResponse(*userOutput))
}
