package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lecterkn/goat_backend/internal/app/handler/request"
	"github.com/lecterkn/goat_backend/internal/app/handler/response"
	"github.com/lecterkn/goat_backend/internal/app/usecase"
	"github.com/lecterkn/goat_backend/internal/app/usecase/input"
)

type AuthorizationHandler struct {
	authUsecase *usecase.AuthorizationUsecase
}

func NewAuthorizationHandler(
	authUsecase *usecase.AuthorizationUsecase,
) *AuthorizationHandler {
	return &AuthorizationHandler{
		authUsecase,
	}
}

// @summary		SignUp
// @description	ユーザーのサインアップを行う
// @tags			user
// @produce		json
// @param			request	body		request.UserSignupRequest	true	"ユーザーログインリクエスト"
// @success		200		{object}	response.UserSignupResponse
// @router			/signup [post]
func (h *AuthorizationHandler) SignUp(ctx echo.Context) error {
	userSignupRequest := request.UserSignupRequest{}
	if err := ctx.Bind(&userSignupRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "invalid request body",
		})
	}
	output, err := h.authUsecase.CreateUser(input.UserCreateInput{
		Username: userSignupRequest.Username,
		Email:    userSignupRequest.Email,
		Password: userSignupRequest.Password,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, response.UserSignupResponse(*output))
}

// @summary		SignIn
// @description	ユーザーのサインインを行う
// @tags			user
// @produce		json
// @param			request	body		request.UserSigninRequest	true	"ユーザーログインリクエスト"
// @success		200		{object}	response.UserSigninResponse
// @router			/signin [post]
func (h *AuthorizationHandler) SignIn(ctx echo.Context) error {
	userSigninRequest := request.UserSigninRequest{}
	if err := ctx.Bind(&userSigninRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "invalid request body",
		})
	}
	output, err := h.authUsecase.LoginUser(input.UserLoginInput{
		Username: userSigninRequest.Username,
		Password: userSigninRequest.Password,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, response.UserSigninResponse(*output))
}
