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
//	@security		BearerAuth
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

//	@summary		GetUsers
//	@description	ユーザーを一覧取得する
//	@tags			user
//	@produce		json
//	@security		BearerAuth
//	@success		200	{object}	response.UserListResponse
//	@router			/users [get]
func (h *UserHandler) GetUsers(ctx echo.Context) error {
	userId, err := uuid.Parse(ctx.Get("userId").(string))
	if err != nil {
		return err
	}
	listOutput, err := h.userUsecase.GetUsers(userId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
	}
	list := []response.UserResponse{}
	for _, outputItem := range listOutput {
		list = append(list, response.UserResponse(outputItem))
	}
	return ctx.JSON(http.StatusOK, response.UserListResponse{
		List: list,
	})
}

//	@summary		EditPermission
//	@description	ユーザーを一覧取得する
//	@tags			user
//	@produce		json
//	@security		BearerAuth
//	@param			userId	path		string								true	"編集対象ユーザーID"
//	@param			request	body		request.UserUpdatePermissionRequest	true	"ユーザーログインリクエスト"
//	@success		200		{object}	response.UserListResponse
//	@router			/users/{userId}/permissions [patch]
func (h *UserHandler) EditPermission(ctx echo.Context) error {
	// 編集者ユーザーID
	userId, err := uuid.Parse(ctx.Get("userId").(string))
	if err != nil {
		return err
	}
	// 編集対象ユーザーID
	targetUserId, err := uuid.Parse(ctx.Param("userId"))
	if err != nil {
		return err
	}
	// リクエスト取得
	updateRequest := request.UserUpdatePermissionRequest{}
	if err := ctx.Bind(&updateRequest); err != nil {
		return err
	}
	// 権限編集
	err = h.userUsecase.EditUserPermission(userId, targetUserId, input.UserUpdatePermissionInput{
		Role: updateRequest.Permission,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
	}
	return ctx.NoContent(http.StatusNoContent)
}
