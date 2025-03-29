package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lecterkn/goat_backend/internal/app/entity"
	"github.com/lecterkn/goat_backend/internal/app/port"
	"github.com/lecterkn/goat_backend/internal/app/repository/redis/model"
	"github.com/redis/go-redis/v9"
)

type TokenRepositoryImpl struct {
	redisClient *redis.Client
}

func NewTokenRepositoryImpl(redisClient *redis.Client) port.TokenRepository {
	return &TokenRepositoryImpl{
		redisClient,
	}
}

// リフレッシュトークンをRedisに追加
func (r *TokenRepositoryImpl) SaveRefreshToken(refreshToken *entity.RefreshTokenEntity) error {
	redisKey := fmt.Sprintf("refreshToken:%s", refreshToken.UserId)
	// json化
	member, err := json.Marshal(model.RedisRefreshToken{
		Id:    refreshToken.Id.String(),
		Token: refreshToken.Token,
	})
	if err != nil {
		return err
	}
	// redisに追加
	return r.redisClient.ZAdd(context.Background(), redisKey, redis.Z{
		Score:  float64(refreshToken.ExpiresIn.UnixMicro()),
		Member: member,
	}).Err()
}

// ユーザーのリフレッシュトークンをRedisから一覧取得
func (r *TokenRepositoryImpl) FindRefreshTokenByUserId(userId uuid.UUID) ([]entity.RefreshTokenEntity, error) {
	redisKey := fmt.Sprintf("refreshToken:%s", userId)
	// 対象ユーザーのリフレッシュトークンをすべて取得
	result, err := r.redisClient.ZRangeWithScores(context.Background(), redisKey, 0, -1).Result()
	if err != nil {
		return nil, err
	}
	var tokenList []entity.RefreshTokenEntity
	for _, token := range result {
		var redisRefreshToken model.RedisRefreshToken
		// Redisから取得した値を構造体化
		err := json.Unmarshal([]byte(token.Member.(string)), &redisRefreshToken)
		if err != nil {
			// ここでエラーが出るのは異常動作のためログ出力した
			fmt.Println(err)
			continue
		}
		// idの文字列をUUIDに変換
		id, err := uuid.Parse(redisRefreshToken.Id)
		if err != nil {
			// ここでエラーが出るのは異常動作のためログ出力した
			fmt.Println(err)
			continue
		}
		// Entity化し、リストに追加
		tokenList = append(tokenList, entity.RefreshTokenEntity{
			Id:        id,
			UserId:    userId,
			Token:     redisRefreshToken.Token,
			ExpiresIn: time.UnixMicro(int64(token.Score)),
		})
	}
	return tokenList, nil
}
