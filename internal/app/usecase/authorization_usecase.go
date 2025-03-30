package usecase

import (
	"context"
	"errors"

	"github.com/lecterkn/goat_backend/internal/app/entity"
	"github.com/lecterkn/goat_backend/internal/app/port"
	"github.com/lecterkn/goat_backend/internal/app/usecase/input"
	"github.com/lecterkn/goat_backend/internal/app/usecase/output"
	"golang.org/x/crypto/bcrypt"
)

type AuthorizationUsecase struct {
	userRepository  port.UserRepository
	tokenRepository port.TokenRepository
}

func NewAuthorizationUsecase(
	userRepository port.UserRepository,
	tokenRepository port.TokenRepository,
) *AuthorizationUsecase {
	return &AuthorizationUsecase{
		userRepository,
		tokenRepository,
	}
}

// ユーザーを新規作成する
func (u *AuthorizationUsecase) CreateUser(cmd input.UserCreateInput) (*output.UserCreateOutput, error) {
	userEntity, err := entity.NewUserEntity(cmd.Username, cmd.Email, cmd.Password, 0)
	if err != nil {
		return nil, err
	}
	err = u.userRepository.Create(context.Background(), userEntity)
	if err != nil {
		return nil, err
	}
	return &output.UserCreateOutput{
		Id:        userEntity.Id,
		Name:      userEntity.Name,
		Email:     userEntity.Email,
		CreatedAt: userEntity.CreatedAt,
		UpdatedAt: userEntity.UpdatedAt,
	}, nil
}

// ユーザーにログインする
func (u *AuthorizationUsecase) LoginUser(cmd input.UserLoginInput) (*output.UserLoginOutput, error) {
	// ユーザー取得
	userEntity, err := u.userRepository.FindByName(context.Background(), cmd.Username)
	if err != nil {
		return nil, err
	}
	// パスワードを確認
	if err := bcrypt.CompareHashAndPassword(userEntity.Password, []byte(cmd.Password)); err != nil {
		return nil, errors.New("invalid password")
	}
	accessTokenEntity, err := entity.NewAccessTokenEntity(userEntity.Id)
	if err != nil {
		return nil, err
	}
	refreshTokenEntity, err := entity.NewRefreshTokenEntity(userEntity.Id)
	if err != nil {
		return nil, err
	}
	// リフレッシュトークンを保存
	err = u.tokenRepository.SaveRefreshToken(refreshTokenEntity)
	if err != nil {
		return nil, err
	}
	return &output.UserLoginOutput{
		AccessToken:  accessTokenEntity.Token,
		RefreshToken: refreshTokenEntity.Token,
	}, nil
}
