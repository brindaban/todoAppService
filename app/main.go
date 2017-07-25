package app

import (
	"os"
	"database/sql"
	"todoApp/todoAppService/config"
	"todoApp/todoAppService/fileReader"
	"todoApp/todoAppService/database"
	"todoApp/todoAppService/errorHandler"
	"todoApp/todoAppService/routers"
)

func main()  {
	context := config.Context{}
	errorLogFilePath := "errorLog"
	errorFile, err := os.OpenFile(errorLogFilePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer errorFile.Close()

	context.ErrorLogFile = errorFile

	dbConfigFilePath := "dbConfgFile"

	if len(os.Args) > 1 {
		dbConfigFilePath = os.Args[1]
	}
	dbConfigDataJson,err := fileReader.ReadJsonFile(dbConfigFilePath,context)
	if err != nil {
		os.Exit(1)
	}
	dbInfo := database.CreateDbInfo(dbConfigDataJson)

	context.Db, err = sql.Open("postgres", dbInfo)

	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile,err)
	}

	context.Db.Ping()

	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile, err)
	}

	defer context.Db.Close()

	routers.HandleRequests(context)

}