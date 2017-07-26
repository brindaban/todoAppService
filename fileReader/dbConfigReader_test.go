package fileReader_test

import (
	"testing"
	"todoApp/todoAppService/fileReader"
	"github.com/stretchr/testify/assert"
	"os"
	"todoApp/todoAppService/routers"
)

func TestJsonObject_IsInOrderReturnsTrueWhenAllConfigsAreAvailable(t *testing.T) {
	testDbConfig := fileReader.JsonObject{DB_USER: "db_user", DB_PASSWORD: "db_password", DB_NAME: "db_name"}
	assert.True(t, testDbConfig.IsInOrder())
}


func TestJsonObject_IsInOrderReturnsFalseWhenAnyConfigIsMissing(t *testing.T) {
	testDbConfig := fileReader.JsonObject{DB_PASSWORD: "db_password", DB_NAME: "db_name"}
	assert.False(t, testDbConfig.IsInOrder())
}

func TestReadJsonFileGivesDbConfigsAsJsonForCorrectFile(t *testing.T) {
	fileName := "testDbConfig";
	errorLogFilePath := "testErrorLog"
	errorFile, err := os.OpenFile(errorLogFilePath, os.O_APPEND|os.O_WRONLY, 0600)
	assert.NoError(t, err)
	defer errorFile.Close()
	testContext := &routers.RouterContext{ErrorLogFile: errorFile}

	dbConfig,err := fileReader.ReadJsonFile(fileName, testContext)

	assert.NoError(t, err)
	assert.Equal(t, "test_db",dbConfig.DB_NAME)
	assert.Equal(t, "test_password",dbConfig.DB_PASSWORD)
	assert.Equal(t, "test_db_user",dbConfig.DB_USER)
}

func TestReadJsonFileGivesErrorForWrongFile(t *testing.T) {
	fileName := "wrongFile";
	errorLogFilePath := "testErrorLog"
	errorFile, err := os.OpenFile(errorLogFilePath, os.O_APPEND|os.O_WRONLY, 0600)
	assert.NoError(t, err)
	defer errorFile.Close()
	testContext := &routers.RouterContext{ErrorLogFile: errorFile}

	_, err = fileReader.ReadJsonFile(fileName, testContext)

	assert.Error(t, err)
}
