package database_test

import (
	"testing"
	"todoApp/todoAppService/fileReader"
	"github.com/stretchr/testify/assert"
	"todoApp/todoAppService/database"
)

func TestCreateDbInfoWhenAllConfigsAreAvailable(t *testing.T) {
	dbConfig := fileReader.JsonObject{DB_USER: "db_user", DB_PASSWORD: "db_password", DB_NAME: "db_name"}
	dbInfo := database.CreateDbInfo(dbConfig)

	assert.Equal(t, "user=db_user password=db_password dbname=db_name sslmode=disable", dbInfo)
}


func TestCreateDbInfoWhenAnyConfigIsMissing(t *testing.T) {
	dbConfig := fileReader.JsonObject{DB_USER: "db_user", DB_PASSWORD: "", DB_NAME: "db_name"}
	dbInfo := database.CreateDbInfo(dbConfig)

	assert.Equal(t, "user=postgres password=postgres dbname=taskmanager sslmode=disable", dbInfo)
}