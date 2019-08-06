package db

import (
	"database/sql"
	"fmt"
	"license-server/db/dbutils"
)

var dbConn *sql.DB
var err error

func startConnection() {
	connStr := dbutils.BuildDBString()
	dbConn, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	if err = dbConn.Ping(); err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println("Connected to Database", connStr)
}

func init() {
	startConnection()
}
