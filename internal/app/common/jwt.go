package common

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func EncodeToken(tokenId, userId uuid.UUID, expiresIn time.Time) (*string, error) {
	claims := jwt.MapClaims{
		"sub": userId.String(),
		"tid": tokenId,
		"exp": expiresIn.Unix(),
	}
	// トークン生成
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := refreshToken.SignedString(getJwtSecretKey())
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
		return []byte(getJwtSecretKey()), nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func getJwtSecretKey() []byte {
	secretKey, ok := os.LookupEnv("JWT_SECRET_KEY")
	if !ok {
		panic("環境変数に\"JWT_SECRE_KEY\"が設定されていません")
	}
	return []byte(secretKey)
}
