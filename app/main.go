package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
	"todoApp/todoAppService/database"
	"todoApp/todoAppService/errorHandler"
	"todoApp/todoAppService/fileReader"
	"todoApp/todoAppService/routers"
)

func main() {
	routerCtx := &routers.RouterContext{}
	errorLogFilePath := "errorLog"
	errorFile, err := os.OpenFile(errorLogFilePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer errorFile.Close()

	routerCtx.ErrorLogFile = errorFile

	dbConfigFilePath := "dbConfgFile"

	if len(os.Args) > 1 {
		dbConfigFilePath = os.Args[1]
	}
	dbConfigDataJson, err := fileReader.ReadJsonFile(dbConfigFilePath, routerCtx)
	if err != nil {
		os.Exit(1)
	}
	dbInfo := database.CreateDbInfo(dbConfigDataJson)

	routerCtx.Db, err = sql.Open("postgres", dbInfo)

	if err != nil {
		errorHandler.ErrorHandler(routerCtx.ErrorLogFile, err)
	}

	routerCtx.Db.Ping()

	if err != nil {
		errorHandler.ErrorHandler(routerCtx.ErrorLogFile, err)
	}

	defer routerCtx.Db.Close()

	routers.HandleRoutes(routerCtx)

}
