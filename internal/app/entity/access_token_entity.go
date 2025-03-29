package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/lecterkn/goat_backend/internal/app/common"
)

const (
	ACCESS_TOKEN_EXP = 60 * 60 * 1 // 1時間
)

type AccessTokenEntity struct {
	UserId    uuid.UUID
	Token     string
	ExpiresIn time.Time
}

func NewAccessTokenEntity(userId uuid.UUID) (*AccessTokenEntity, error) {
	// トークン期限設定
	tokenExp := time.Now().Add(time.Second * time.Duration(ACCESS_TOKEN_EXP))
	// トークンID生成
	tokenId := uuid.New()
	// トークンの中身
	token, err := common.EncodeToken(tokenId, userId, tokenExp)
	if err != nil {
		return nil, err
	}
	return &AccessTokenEntity{
		UserId:    userId,
		Token:     *token,
		ExpiresIn: tokenExp,
	}, nil
}
