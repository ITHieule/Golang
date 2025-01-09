package routes

import (
	"Golang/controllers"
	"Golang/handlers"

	"github.com/gorilla/mux"
)

// RegisterTodoRoutes định nghĩa các route cho todo
func RegisterTodoRoutes(r *mux.Router) {
	r.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
	r.HandleFunc("/todos", handlers.CreateTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", handlers.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", handlers.DeleteTodo).Methods("DELETE")
	// Đăng ký route cho login
	r.HandleFunc("/api/login", controllers.LoginHandler).Methods("POST")
}

// SetupRoutes thiết lập tất cả các route API
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	RegisterTodoRoutes(router)
	return router
}
