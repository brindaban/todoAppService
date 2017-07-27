package handlers

import (
	"database/sql"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"strings"
	"strconv"
)

const (
	dbSelectQuery = "select id,description,priority,finished from tasks;"
	dbInsertQuery = "insert into tasks(description,priority,finished) VALUES($1,$2,$3);"
	dbDeleteQuery = "delete from tasks where id = $1;"
	dbUpdateQuery = "update tasks set description = $1, priority = $2, finished = $3 where id = $4;"
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

func AddTodo(db *sql.DB) http.HandlerFunc{
	return func(res http.ResponseWriter, req *http.Request) {
		body,err := ioutil.ReadAll(req.Body)

		if err != nil {
			fmt.Println("Failed to know request body")
			res.WriteHeader(http.StatusBadRequest)
		}

		var todoDTO TodoContent
		err = json.Unmarshal(body, &todoDTO)
		if err != nil {
			fmt.Println("failed to unmarshal")
		}

		_,err = db.Exec(dbInsertQuery, todoDTO.Description, todoDTO.Priority, todoDTO.Finished)
		if err != nil {
			fmt.Println("Failed to insert todo",err)
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusCreated)

	}
}

func DeleteTodo(db *sql.DB) http.HandlerFunc{
	return func(res http.ResponseWriter, req *http.Request) {
		reqTodoId,err := strconv.Atoi(strings.Split(req.URL.String(),"/")[3])
		if err != nil {
			fmt.Println("Failed to get the id from url",err)
		}

		_,err = db.Exec(dbDeleteQuery, reqTodoId)
		if err != nil {
			fmt.Println("Failed to delete todo",err)
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusAccepted)
	}
}

func UpdateTodo(db *sql.DB) http.HandlerFunc{
	return func(res http.ResponseWriter, req *http.Request) {
		body,err := ioutil.ReadAll(req.Body)

		if err != nil {
			fmt.Println("Failed to know request body",err)
			res.WriteHeader(http.StatusBadRequest)
		}

		var todoDTO TodoContent
		err = json.Unmarshal(body, &todoDTO)
		if err != nil {
			fmt.Println("failed to unmarshal",err)
		}
		reqTodoId,err := strconv.Atoi(strings.Split(req.URL.String(),"/")[3])
		if err != nil {
			fmt.Println("Failed to get the id from url",err)
		}

		todoDTO.ID = reqTodoId

		_,err = db.Exec(dbUpdateQuery, todoDTO.Description, todoDTO.Priority, todoDTO.Finished, todoDTO.ID)

		if err != nil {
			fmt.Println("Failed to update todo",err)
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusCreated)
	}
}