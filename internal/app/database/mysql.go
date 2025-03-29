package database

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type TxKeyType struct{}

var TxKey = TxKeyType{}

func GetMySQLConnection() *sqlx.DB {
	db, err := sqlx.Connect("mysql", getMySQLUrl())
	if err != nil {
		panic(err)
	}
	return db
}

func getMySQLUrl() string {
	url, ok := os.LookupEnv("MYSQL_DSN")
	if !ok {
		panic("環境変数に\"MYSQL_DSN\"が設定されていません")
	}
	return url
}
