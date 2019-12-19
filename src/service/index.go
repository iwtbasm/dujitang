package service

import (
	"database/sql"
	"dujitang/src/config"
	"dujitang/src/model"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", config.SQLDB)
	if err != nil {
		panic(err)
	}

	return
}

// 毒鸡汤
func Index() (*model.Soul, error) {
	var soul model.Soul
	err := db.QueryRow("select * from soul order by rand( ) limit 1").Scan(&soul.ID, &soul.Title, &soul.Hits)
	config.CheckErr(err)

	return &soul, nil
}
