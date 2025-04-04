package entity

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type MyListEntity struct {
	UserId  uuid.UUID
	Stories []*ScoredStoryEntity
}

func NewMyListEntity(userId uuid.UUID) *MyListEntity {
	return &MyListEntity{
		UserId:  userId,
		Stories: []*ScoredStoryEntity{},
	}
}

// マイリストにストーリーを追加
func (e *MyListEntity) Add(storyEntity *StoryEntity, score int) error {
	// 重複確認
	for _, mylistItem := range e.Stories {
		if mylistItem.Story.Id == storyEntity.Id {
			return errors.New("the story is already added")
		}
	}
	// 作成
	scoredStoryEntity, err := NewScoredStoryEntity(*storyEntity, score)
	if err != nil {
		return err
	}
	// 追加
	e.Stories = append(e.Stories, scoredStoryEntity)
	return nil
}

type ScoredStoryEntity struct {
	Story StoryEntity
	Score int `validate:"required,min=1,max=10"`
}

func NewScoredStoryEntity(story StoryEntity, score int) (*ScoredStoryEntity, error) {
	scoredStoryEntity := ScoredStoryEntity{
		Story: story,
		Score: score,
	}
	if err := scoredStoryEntity.isValid(); err != nil {
		return nil, err
	}
	return &scoredStoryEntity, nil
}

func (e *ScoredStoryEntity) UpdateScore(score int) error {
	oldScore := e.Score
	e.Score = score
	err := e.isValid()
	if err != nil {
		e.Score = oldScore
	}
	return err
}

func (e *ScoredStoryEntity) isValid() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(e)
}
