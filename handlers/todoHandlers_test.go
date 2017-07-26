package handlers_test

import (
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"net/http"
	"net/http/httptest"

	"todoApp/todoAppService/handlers"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestGetAllTodo(t *testing.T) {
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
	fmt.Println(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, w.Body.String())

}