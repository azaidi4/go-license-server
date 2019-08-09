package db

import (
	"database/sql"
	"license-server/db/dbutils"
	"license-server/utils/logger"

	// Database driver for Postgres
	_ "github.com/lib/pq"
)

var DBConn *sql.DB
var err error

func startConnection() {
	connStr := dbutils.BuildDBString()
	DBConn, err = sql.Open("postgres", connStr)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	if err = DBConn.Ping(); err != nil {
		logger.Error.Println(err)
		return
	}
	logger.Info.Println("Connected to Database", connStr)
}

func initTable() {
	init, err := DBConn.Prepare(
		`CREATE TABLE IF NOT EXISTS dc_licenses(
      id SERIAL PRIMARY KEY,
      license_expiry DATE NOT NULL,
      license_user_limit INT NOT NULL,
      license_key VARCHAR NOT NULL,
      license_org VARCHAR NOT NULL UNIQUE
    );`)
	defer init.Close()
	if err != nil {
		logger.Error.Println(err)
		return
	}
	if _, err := init.Query(); err != nil {
		logger.Error.Println(err)
		return
	}
	logger.Info.Println("Created table dc_licenses")
}

func init() {
	startConnection()
	initTable()
}
