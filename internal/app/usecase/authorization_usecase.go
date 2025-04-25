package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/lecterkn/goat_backend/internal/app/common"
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
func (u *AuthorizationUsecase) CreateUser(cmd input.UserCreateInput) (*output.UserOutput, error) {
	userEntity, err := entity.NewUserEntity(cmd.Username, cmd.Email, cmd.Password, 0)
	if err != nil {
		return nil, err
	}
	err = u.userRepository.Create(context.Background(), userEntity)
	if err != nil {
		return nil, err
	}
	return &output.UserOutput{
		Id:        userEntity.Id,
		Name:      userEntity.Name,
		Email:     userEntity.Email,
		Role:      uint8(userEntity.Role.Permission),
		RoleName:  userEntity.Role.GetPermission(),
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

func (u *AuthorizationUsecase) RefreshAccessToken(cmd input.RefreshInput) (*output.RefreshOutput, error) {
	// トークンをデコード
	claims, err := common.DecodeToken(cmd.RefreshToken)
	if err != nil {
		return nil, err
	}
	// sub取得
	sub, err := claims.GetSubject()
	if err != nil {
		return nil, err
	}
	// UUIDに変換
	userId, err := uuid.Parse(sub)
	if err != nil {
		return nil, err
	}
	// 対象ユーザーのリフレッシュトークン一覧取得
	refreshTokenEntities, err := u.tokenRepository.FindRefreshTokenByUserId(userId)
	if err != nil {
		return nil, errors.New("ユーザーの認証情報が見つかりませんでした")
	}
	// トークン確認
	for _, refreshTokenEntity := range refreshTokenEntities {
		// 期限確認
		if !refreshTokenEntity.ExpiresIn.After(time.Now()) {
			continue
		}
		// 一致した場合
		if refreshTokenEntity.Token == cmd.RefreshToken {
			// アクセストークン発行
			accessTokenEntity, err := entity.NewAccessTokenEntity(userId)
			if err != nil {
				return nil, err
			}
			return &output.RefreshOutput{
				AccessToken: accessTokenEntity.Token,
				ExpiresIn:   accessTokenEntity.ExpiresIn,
			}, nil
		}
	}
	// 一致するリフレッシュトークンが存在しない場合
	return nil, errors.New("不正なリフレッシュトークンです")
}
