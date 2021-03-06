package routers

import (
	"todoApp/todoAppService/constants"
	"todoApp/todoAppService/handlers"

	"os"
	"github.com/gorilla/mux"
	"net/http"
	"database/sql"
	"fmt"
)


type RouterContext struct {
	ErrorLogFile *os.File
	Db *sql.DB
}

type Route struct {
	Path                string
	Handler             http.HandlerFunc
	Methods             []string
}

func registerRoutes(r *mux.Router, context *RouterContext) {
	routes := getAPIRoutes(context)
	for _, route := range routes {
		r.HandleFunc(route.Path, route.Handler).Methods(route.Methods...)
	}
}
func getAPIRoutes(context *RouterContext) []Route {
	db := context.Db

	return []Route{
		{
			Path: constants.GetAllTodoPath,
			Handler: handlers.GetAllTodo(db),
			Methods: []string {"GET"},
		},
		{
			Path: constants.CreateTodoPath,
			Handler: handlers.AddTodo(db),
			Methods: []string {"POST"},
		},
		{
			Path: constants.DeleteTodoPath,
			Handler: handlers.DeleteTodo(db),
			Methods: []string {"DELETE"},
		},
		{
			Path: constants.UpdateTodoPath,
			Handler: handlers.UpdateTodo(db),
			Methods: []string {"PATCH"},
		},

	}
}


func HandleRoutes(context *RouterContext) {
	r := mux.NewRouter()
	registerRoutes(r, context)

	srv := &http.Server{
		Addr:    ":8080", // Normally ":443"
		Handler: r,
	}
	fmt.Println("Server starting at port",srv.Addr)
	srv.ListenAndServe()
}