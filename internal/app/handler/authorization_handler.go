package handler

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/lecterkn/goat_backend/internal/app/common"
	"github.com/lecterkn/goat_backend/internal/app/handler/request"
	"github.com/lecterkn/goat_backend/internal/app/handler/response"
	"github.com/lecterkn/goat_backend/internal/app/usecase"
	"github.com/lecterkn/goat_backend/internal/app/usecase/input"
)

const (
	AUTHORIZATION_HEADER      = "Authorization"
	AUTHORIZATION_PREFIX      = "Bearer "
	REFRESH_TOKEN_HEADER_NAME = "x-refresh-token"
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

//	@summary		SignUp
//	@description	ユーザーのサインアップを行う
//	@tags			auth
//	@produce		json
//	@security		BearerAuth
//	@param			request	body		request.UserSignupRequest	true	"ユーザーログインリクエスト"
//	@success		200		{object}	response.UserResponse
//	@router			/signup [get]
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
	return ctx.JSON(http.StatusOK, response.UserResponse(*output))
}

//	@summary		SignIn
//	@description	ユーザーのサインインを行う
//	@tags			auth
//	@produce		json
//	@param			request	body		request.UserSigninRequest	true	"ユーザーログインリクエスト"
//	@success		200		{object}	response.UserSigninResponse
//	@router			/signin [post]
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

//	@summary		Refresh
//	@description	アクセストークンをリフレッシュする
//	@tags			auth
//	@produce		json
//	@param			x-refresh-token	header		string	true	"リフレッシュトークン"
//	@success		200				{object}	response.RefreshResponse
//	@router			/refresh [post]
func (h *AuthorizationHandler) Refresh(ctx echo.Context) error {
	// ヘッダーからリフレッシュトークン取得
	refreshToken := ctx.Request().Header.Get(REFRESH_TOKEN_HEADER_NAME)
	// リフレッシュ処理
	output, err := h.authUsecase.RefreshAccessToken(input.RefreshInput{
		RefreshToken: refreshToken,
	})
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{
			Message: err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, response.RefreshResponse(*output))
}

// JWT認証を行うミドルウェア
func (h *AuthorizationHandler) Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// ヘッダーからアクセストークンを取得
		header := ctx.Request().Header.Get(AUTHORIZATION_HEADER)
		// トークンの形式を確認
		if !strings.HasPrefix(header, AUTHORIZATION_PREFIX) {
			return ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{
				Message: "不正な認証ヘッダー",
			})
		}
		// トークン部分のみ取得
		token := strings.TrimPrefix(header, AUTHORIZATION_PREFIX)
		// トークンを検証・復号化
		claims, err := common.DecodeToken(token)
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{
				Message: "トークンの復号化に失敗しました",
			})
		}
		// トークンからID取得
		sub, err := claims.GetSubject()
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{
				Message: "トークンからユーザーIDを取得できませんでした",
			})
		}
		// コンテキストにユーザーIDを設定
		ctx.Set("userId", sub)
		// エンドポイントに処理を引き渡す
		return next(ctx)
	}
}
