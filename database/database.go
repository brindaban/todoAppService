package database

import (
	"fmt"
	"todoApp/todoAppService/fileReader"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "taskmanager"
)

func CreateDbInfo(dbConfig fileReader.JsonObject) string {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)

	if (dbConfig.IsInOrder()){
		dbinfo = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
			dbConfig.DB_USER, dbConfig.DB_PASSWORD, dbConfig.DB_NAME)
	}

	return dbinfo
}
