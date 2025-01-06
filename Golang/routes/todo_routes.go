package routes

import (
	"Golang/handlers"
	"github.com/gorilla/mux"
)

// SetupRoutes định nghĩa các route API
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Định nghĩa các route API
	router.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
	router.HandleFunc("/todos", handlers.CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", handlers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", handlers.DeleteTodo).Methods("DELETE")

	return router
}
