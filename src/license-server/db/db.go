package db

import (
	"database/sql"
	"fmt"
	"license-server/db/dbutils"

	_ "github.com/lib/pq"
)

var DBConn *sql.DB
var err error

func startConnection() {
	connStr := dbutils.BuildDBString()
	DBConn, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	if err = DBConn.Ping(); err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println("Connected to Database", connStr)
}

func initTable() {
	init, err := DBConn.Prepare(
		`CREATE TABLE IF NOT EXISTS dc_licenses(
      id SERIAL PRIMARY KEY,
      license_expiry DATE NOT NULL,
      license_user_limit INT NOT NULL,
      license_key BIGINT NOT NULL,
      license_org VARCHAR NOT NULL
    );
    `)
	defer init.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	if _, err := init.Query(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Created table dc_licenses")
}

func init() {
	startConnection()
	initTable()
}
