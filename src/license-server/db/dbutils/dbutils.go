package dbutils

import (
	"fmt"
	"license-server/env"
)

// BuildDBURL returns [schema]://[user[:password]@][host][:port][/dbname][?param1=value1&...]
func BuildDBURL(schema string) string {
	db := env.Config.Database
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", schema, db.User, db.Password, db.Host, db.Port, db.Name)
}

// BuildDBString returns DB connection string
func BuildDBString() string {
	db := env.Config.Database
	return fmt.Sprintf("user=%s dbname=%s host=%s port=%s sslmode=disable",
		db.User, db.Name, db.Host, db.Port)
}
