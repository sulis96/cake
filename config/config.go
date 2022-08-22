package config

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func MySqlDatabase() (*sql.DB, error) {
	dbConfig := fmt.Sprintf("%v:%v@/%v?parseTime=true", "root", "", "cake")
	db, err := sql.Open("mysql", dbConfig)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(50)
	db.SetConnMaxIdleTime(time.Second * 60)

	return db, nil
}
