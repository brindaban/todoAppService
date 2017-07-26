package handlers

import (
	"database/sql"
	"os"
	"net/http"
	"io"
)

func GetAllTodo(db *sql.DB, loggingFile *os.File) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func ServeIndex() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "hello, world!\n")
	}
}
