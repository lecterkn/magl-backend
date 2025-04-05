package mysql

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lecterkn/goat_backend/internal/app/entity"
	"github.com/lecterkn/goat_backend/internal/app/port"
	"github.com/lecterkn/goat_backend/internal/app/repository/mysql/model"
)

type MyListRepositoryImpl struct {
	database *sqlx.DB
}

func NewMyListRepositoryImpl(database *sqlx.DB) port.MyListRepository {
	return &MyListRepositoryImpl{
		database,
	}
}

func (r *MyListRepositoryImpl) Save(ctx context.Context, mylistEntity *entity.MyListEntity) error {
	return RunInTx(ctx, r.database, func(tx *sqlx.Tx) error {
		// 全て削除
		query := `
            DELETE FROM mylists
            WHERE mylists.user_id = :userId
        `
		queryMap := map[string]any{
			"userId": mylistEntity.UserId[:],
		}
		_, err := tx.NamedExec(query, queryMap)
		if err != nil {
			return err
		}
		if len(mylistEntity.Stories) == 0 {
			return nil
		}
		// 全て挿入
		query = `
            INSERT INTO mylists(user_id, story_id, score)
            VALUES 
        `
		for i, scoredStory := range mylistEntity.Stories {
			query += fmt.Sprintf("(:userId, :storyId%d, :score%d)", i, i)
			queryMap[fmt.Sprintf("storyId%d", i)] = scoredStory.Story.Id[:]
			queryMap[fmt.Sprintf("score%d", i)] = scoredStory.Score
			if i < len(mylistEntity.Stories)-1 {
				query += ","
			}
		}
		_, err = tx.NamedExec(query, queryMap)
		if err != nil {
			return err
		}
		return nil
	})
}

func (r *MyListRepositoryImpl) FindByUserId(ctx context.Context, userId uuid.UUID) (*entity.MyListEntity, error) {
	query := `
        SELECT 
            mylists.story_id, mylists.score,
            stories.title AS story_title, stories.episode AS story_episode, stories.description AS story_desc, stories.image_url AS story_image_url, stories.created_at AS story_created_at, stories.updated_at AS story_updated_at,
            categories.id As category_id, categories.name AS category_name, categories.description AS category_desc, categories.image_url AS category_image_url, 
            categories.created_at AS category_created_at, categories.updated_at AS category_updated_at
        FROM mylists
        JOIN stories 
            ON stories.id = mylists.story_id 
        JOIN categories 
            ON categories.id = stories.category_id
        WHERE mylists.user_id = ?
    `
	mylistModels := []model.MyListModel{}
	err := RunInTx(ctx, r.database, func(tx *sqlx.Tx) error {
		return tx.Select(&mylistModels, query, userId[:])
	})
	if err != nil {
		return nil, err
	}
	return r.toEntity(userId, mylistModels)
}

func (r *MyListRepositoryImpl) toEntity(userId uuid.UUID, mylistModels []model.MyListModel) (*entity.MyListEntity, error) {
	scoredStories := []*entity.ScoredStoryEntity{}
	for _, mylistModel := range mylistModels {
		storyId, err := uuid.FromBytes(mylistModel.StoryId)
		if err != nil {
			return nil, err
		}
		var storyImageUrl *string = nil
		if mylistModel.StoryImageUrl.Valid {
			storyImageUrl = &mylistModel.StoryImageUrl.String
		}
		categoryId, err := uuid.FromBytes(mylistModel.CategoryId)
		if err != nil {
			return nil, err
		}
		var categoryImageUrl *string = nil
		if mylistModel.CategoryImageUrl.Valid {
			categoryImageUrl = &mylistModel.CategoryImageUrl.String
		}
		scoredStories = append(scoredStories, &entity.ScoredStoryEntity{
			Score: mylistModel.Score,
			Story: entity.StoryEntity{
				Id: storyId,
				Category: entity.CategoryEntity{
					Id:          categoryId,
					Name:        mylistModel.CategoryName,
					Description: mylistModel.CategoryDescription,
					ImageUrl:    categoryImageUrl,
					CreatedAt:   mylistModel.CategoryCreatedAt,
					UpdatedAt:   mylistModel.CategoryUpdatedAt,
				},
				Title:       mylistModel.StoryTitle,
				Episode:     mylistModel.StoryEpisode,
				Description: mylistModel.StoryDescription,
				ImageUrl:    storyImageUrl,
				CreatedAt:   mylistModel.StoryCreatedAt,
				UpdatedAt:   mylistModel.StoryUpdatedAt,
			},
		})
	}
	return &entity.MyListEntity{
		UserId:  userId,
		Stories: scoredStories,
	}, nil
}
