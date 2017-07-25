package fileReader_test

import (
	"testing"
	"todoApp/todoAppService/fileReader"
	"github.com/stretchr/testify/assert"
	"todoApp/todoAppService/config"
	"os"
)

func generateContextObject()(config.Context){
	context := config.Context{}
	errorLogFilePath := "../errorLog"
	errorFile, err := os.OpenFile(errorLogFilePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer errorFile.Close()
	context.ErrorLogFile = errorFile
	return context
}

func TestJsonObjectIsInOrderReturnTrue(t *testing.T)  {
	jsonObj := fileReader.JsonObject{"Brinda","Password","dbName","db_schema"}
	assert.True(t, jsonObj.IsInOrder())
}

func TestJsonObjectIsInOrderReturnFalseWhenDbNameIsEmpty(t *testing.T)  {
	jsonObj := fileReader.JsonObject{"Brinda","Password","","db_schema"}
	assert.False(t, jsonObj.IsInOrder())
}

func TestReadJsonFileGivesJsonObjectForCorrectFile(t *testing.T)  {
	testConfFile := "dbConfigForTest";
	jsonObj,err := fileReader.ReadJsonFile(testConfFile, config.Context{});
	assert.NoError(t, err);
	assert.Equal(t, "postgres_test", jsonObj.DB_USER)
}

func TestReadJsonFileThrowsErrorForIncorrectFile(t *testing.T)  {
	testConfFile := "incorrectFile";
	_,err := fileReader.ReadJsonFile(testConfFile, generateContextObject());
	assert.Error(t, err);
}
