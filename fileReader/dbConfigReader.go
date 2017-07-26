package fileReader

import (
	"io/ioutil"
	"encoding/json"
	"todoApp/todoAppService/errorHandler"
	"todoApp/todoAppService/routers"
)

type JsonObject struct {
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

func (db *JsonObject) IsInOrder() bool {
	return db.DB_NAME != "" && db.DB_USER != "" && db.DB_PASSWORD != ""
}

func ReadJsonFile(fileName string, context *routers.RouterContext) (JsonObject, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile, err)
		return JsonObject{}, err
	}
	var jsonType JsonObject
	json.Unmarshal(file, &jsonType)
	return jsonType, nil
}
