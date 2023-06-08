package db

import (
	"database/sql"
	"time"

	"github.com/andrefsilveira1/LoadEnv"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	cfg := mysql.Config{
		User:      LoadEnv.LoadEnv("DB_USER"),
		Passwd:    LoadEnv.LoadEnv("DB_PASS"),
		Net:       "tcp",
		Addr:      "172.17.0.2:3306",
		DBName:    "todo",
		ParseTime: true,
		Loc:       &time.Location{},
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		panic(err.Error())
	}

	return db, nil

}
