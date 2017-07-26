package handlers

import (
	"database/sql"
	"net/http"
	"io"
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

func GetAllTodo(db *sql.DB) http.HandlerFunc{
	return func(res http.ResponseWriter, req *http.Request) {
		rows, err := db.Query(dbSelectQuery)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
		}

		dbData := []TodoContent{}
		if (rows != nil) {
			for rows.Next() {
				var r TodoContent
				rows.Scan(&r.ID, &r.Description, &r.Priority,&r.Finished)
				dbData = append(dbData, r)
			}
		}
		data,_ := json.Marshal(dbData)

		res.WriteHeader(http.StatusOK)
		res.Write(data)
	}
}

func ServeIndex() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "hello, world!\n")
	}
}
