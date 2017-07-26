package database

import (
	"fmt"
	"todoApp/todoAppService/fileReader"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "taskmanager"
	DB_SCHEMA   = "tasks"
)

func CreateDbInfo(dbConfig fileReader.JsonObject) string {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable SEARCH_PATH=%s ",
		DB_USER, DB_PASSWORD, DB_NAME, DB_SCHEMA)

	if (dbConfig.IsInOrder()){
		dbinfo = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable SEARCH_PATH=%s ",
			dbConfig.DB_USER, dbConfig.DB_PASSWORD, dbConfig.DB_NAME, dbConfig.DB_SCHEMA)
	}

	return dbinfo
}
