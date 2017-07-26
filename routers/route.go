package routers

import (
	"todoApp/todoAppService/constants"
	"todoApp/todoAppService/handlers"

	"os"
	"github.com/gorilla/mux"
	"net/http"
	"database/sql"
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
	errorLogFile := context.ErrorLogFile

	return []Route{
		{
			Path: constants.IndexPath,
			Handler: handlers.ServeIndex(),
			Methods: []string{"GET"},
		},
		{
			Path: constants.GetAllTodoPath,
			Handler: handlers.GetAllTodo(db, errorLogFile),
			Methods: []string {"GET"},
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
	srv.ListenAndServe()
}