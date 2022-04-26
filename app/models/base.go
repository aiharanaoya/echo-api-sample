package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
var err error

// DB接続
func ConnectDb() {
	// 第二引数: userName:password@(host:port)/dbName
	Db, err = sql.Open("mysql", "root:root_password@tcp(db:3306)/echo_api_sample?parseTime=true")

	if err != nil {
		fmt.Println(err.Error())
	}

	// 接続が終了したらクローズする
	defer Db.Close()

	// DB接続テスト
	err = Db.Ping()

	if err != nil {
		fmt.Println("DB接続失敗", err.Error())
	} else {
		fmt.Println("DB接続成功")
	}
}
