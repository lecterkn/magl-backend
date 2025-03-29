package mysql

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/lecterkn/goat_backend/internal/app/database"
)

// RunInTx
// db データベース
// ctx トランザクションが存在しない場合は新規トランザクションをコミットまで行う
// txFn 処理内容
func RunInTx(ctx context.Context, db *sqlx.DB, txFn func(tx *sqlx.Tx) error) error {
	tx := GetTx(ctx)
	if tx == nil {
		// トランザクション開始
		tx, err := db.Beginx()
		if err != nil {
			return err
		}
		// 処理実行
		err = txFn(tx)
		// エラーの場合はロールバック
		if err != nil {
			tx.Rollback()
			return err
		}
		// コミット
		return tx.Commit()
	}
	return txFn(tx)
}

func GetTx(ctx context.Context) *sqlx.Tx {
	tx, ok := ctx.Value(database.TxKey).(*sqlx.Tx)
	if !ok {
		return nil
	}
	return tx
}
