package handlers_test

import (
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"net/http"
	"net/http/httptest"

	"todoApp/todoAppService/handlers"
	"github.com/stretchr/testify/assert"
	"errors"
	"bytes"
)

func TestGetAllTodoSuccess(t *testing.T) {
	db, mock, _ := sqlmock.New()

	r, err := http.NewRequest("GET", "/api/todo", nil)
	assert.NoError(t, err, "failed to make a GET request")

	w := httptest.NewRecorder()

	columnOrder := []string{"id", "description", "priority", "finished"}
	mockRows := sqlmock.NewRows(columnOrder).
		AddRow(1, "Description1", "High", true).
		AddRow(1, "Description2", "Low", false)
	mock.ExpectQuery("select id,description,priority,finished from tasks;").
		WillReturnRows(mockRows)

	handlers.GetAllTodo(db)(w, r)
	expected := "[{\"ID\":1,\"Description\":\"Description1\",\"Priority\":\"High\",\"Finished\":true},{\"ID\":1,\"Description\":\"Description2\",\"Priority\":\"Low\",\"Finished\":false}]"
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, w.Body.String())
}

func TestGetAllTodoFailsForDbError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	r, err := http.NewRequest("GET", "/api/todo", nil)
	assert.NoError(t, err, "failed to make a GET request")

	w := httptest.NewRecorder()

	mock.ExpectQuery("select id,description,priority,finished from tasks;").WillReturnError(errors.New("DB Error"))

	handlers.GetAllTodo(db)(w, r)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestAddTodoSuccess(t *testing.T) {
	db, mock, _ := sqlmock.New()

	todo := []byte(`{
						"Description" : "New Todo",
						"Priority" : "HIGH" ,
						"Finished" : true
					}`)


	r, err := http.NewRequest("POST", "/api/todo/new",bytes.NewBuffer(todo))

	assert.NoError(t, err, "failed to make a GET request")

	w := httptest.NewRecorder()

	mock.ExpectExec("insert into tasks").
		WithArgs("New Todo", "HIGH", true).WillReturnResult(sqlmock.NewResult(1,1))

	handlers.AddTodo(db)(w, r)
	assert.Equal(t, http.StatusCreated, w.Code)
}
