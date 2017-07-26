package handlers

import (
	"database/sql"
	"os"
	"net/http"
	"io"
	"todoApp/todoAppService/errorHandler"
	"fmt"
	"encoding/json"
)

const (
	dbSelectQuery string = "select id,description,priority,finished from tasks;"
)

type TodoContent struct{
	ID int
	Description string
	Priority string
	Finished bool
}

func GetAllTodo(db *sql.DB, loggingFile *os.File) http.HandlerFunc{
	return func(res http.ResponseWriter, req *http.Request) {
		rows, err := db.Query(dbSelectQuery)
		if err != nil {
			errorHandler.ErrorHandler(loggingFile, err)
		}

		dbData := []TodoContent{}
		if (rows != nil) {
			for rows.Next() {
				var r TodoContent
				rows.Scan(&r.ID, &r.Description, &r.Priority,&r.Finished)
				dbData = append(dbData, r)
			}
		}
		fmt.Println(dbData,"Here I am")
		data,_ := json.Marshal(dbData)
		res.Write(data)
	}
}

func ServeIndex() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "hello, world!\n")
	}
}
