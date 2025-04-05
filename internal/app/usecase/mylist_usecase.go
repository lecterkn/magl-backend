package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/lecterkn/goat_backend/internal/app/port"
	"github.com/lecterkn/goat_backend/internal/app/usecase/input"
	"github.com/lecterkn/goat_backend/internal/app/usecase/output"
)

type MyListUsecase struct {
	storyRepository  port.StoryRepository
	mylistRepository port.MyListRepository
	txProvider       port.TransactionProvider
}

func NewMyListUsecase(
	storyRepository port.StoryRepository,
	mylistRepository port.MyListRepository,
	txProvider port.TransactionProvider,
) *MyListUsecase {
	return &MyListUsecase{
		storyRepository,
		mylistRepository,
		txProvider,
	}
}

// マイリストにストーリーをスコアを付けて保存
func (u *MyListUsecase) AddStoryToMyList(userId uuid.UUID, cmd input.MyListAddInput) error {
	// トランザクション開始
	return u.txProvider.Transact(func(ctx context.Context) error {
		// マイリスト取得
		mylistEntity, err := u.mylistRepository.FindByUserId(ctx, userId)
		if err != nil {
			return err
		}
		// ストーリー取得
		storyEntity, err := u.storyRepository.FindById(ctx, cmd.StoryId)
		if err != nil {
			return err
		}
		// マイリストに追加
		err = mylistEntity.Add(storyEntity, cmd.Score)
		if err != nil {
			return err
		}
		// 保存
		return u.mylistRepository.Save(ctx, mylistEntity)
	})
}

// マイリストに登録済みストーリーのスコアを更新
func (u *MyListUsecase) UpdateScore(userId uuid.UUID, cmd input.MyListUpdateInput) error {
	// トランザクション開始
	return u.txProvider.Transact(func(ctx context.Context) error {
		// マイリスト取得
		mylistEntity, err := u.mylistRepository.FindByUserId(ctx, userId)
		if err != nil {
			return err
		}
		for _, story := range mylistEntity.Stories {
			// 対象ストーリー検索
			if story.Story.Id == cmd.StoryId {
				// スコア更新
				err := story.UpdateScore(cmd.Score)
				if err != nil {
					return err
				}
				// マイリスト保存
				return u.mylistRepository.Save(ctx, mylistEntity)
			}
		}
		return errors.New("the story is not in mylist")
	})
}

// マイリストをユーザーIDから取得
func (u *MyListUsecase) GetMylist(userId uuid.UUID) (*output.MyListOutput, error) {
	mylistEntity, err := u.mylistRepository.FindByUserId(context.Background(), userId)
	if err != nil {
		return nil, err
	}
	outputList := []output.MyListItemOutput{}
	for _, scoredStoryEntity := range mylistEntity.Stories {
		outputList = append(outputList, output.MyListItemOutput{
			StoryId:      scoredStoryEntity.Story.Id,
			CategoryId:   scoredStoryEntity.Story.Category.Id,
			CategoryName: scoredStoryEntity.Story.Category.Name,
			Title:        scoredStoryEntity.Story.Title,
			Episode:      scoredStoryEntity.Story.Episode,
			Description:  scoredStoryEntity.Story.Description,
			ImageUrl:     scoredStoryEntity.Story.ImageUrl,
			Score:        scoredStoryEntity.Score,
		})
	}
	return &output.MyListOutput{
		UserId:  mylistEntity.UserId,
		Stories: outputList,
	}, nil
}

// マイリストからストーリーを削除
func (u *MyListUsecase) RemoveFromList(userId uuid.UUID, storyId uuid.UUID) error {
	// トランザクション開始
	return u.txProvider.Transact(func(ctx context.Context) error {
		// マイリスト取得
		mylistEntity, err := u.mylistRepository.FindByUserId(ctx, userId)
		if err != nil {
			return err
		}
		// ストーリーを削除
		err = mylistEntity.Remove(storyId)
		if err != nil {
			return err
		}
		// 保存
		return u.mylistRepository.Save(ctx, mylistEntity)
	})
}
