package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/lecterkn/goat_backend/internal/app/common"
)

const (
	REFRESH_TOKEN_EXP = 60 * 60 * 24 * 180 // 180日
)

type RefreshTokenEntity struct {
	Id        uuid.UUID
	UserId    uuid.UUID
	Token     string
	ExpiresIn time.Time
}

func NewRefreshTokenEntity(userId uuid.UUID) (*RefreshTokenEntity, error) {
	// トークン期限設定
	tokenExp := time.Now().Add(time.Second * time.Duration(REFRESH_TOKEN_EXP))
	// トークンID生成
	tokenId := uuid.New()
	// トークンの中身
	token, err := common.EncodeToken(tokenId, userId, tokenExp)
	if err != nil {
		return nil, errors.New("リフレッシュトークンの生成に失敗しました")
	}
	return &RefreshTokenEntity{
		Id:        tokenId,
		UserId:    userId,
		Token:     *token,
		ExpiresIn: tokenExp,
	}, nil
}
