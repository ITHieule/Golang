package handlers

import (
	"Golang/database"
	"Golang/models"
	"encoding/json"

	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// GetTodos trả về danh sách công việc
func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	todos, err := database.GetTodos()
	if err != nil {
		http.Error(w, "Lỗi khi lấy dữ liệu", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todos)
}

// CreateTodo thêm một công việc mới
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newTodo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		http.Error(w, "Dữ liệu không hợp lệ", http.StatusBadRequest)
		return
	}

	id, err := database.CreateTodo(newTodo)
	if err != nil {
		http.Error(w, "Không thể thêm công việc", http.StatusInternalServerError)
		return
	}

	newTodo.ID = int(id)
	json.NewEncoder(w).Encode(newTodo)
}

// UpdateTodo cập nhật trạng thái công việc
// Hàm để xử lý API PUT cho UpdateTodo
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	// Parse URL parameters, assume id is part of the URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Parse body data
	var updatedTodo models.Todo
	err = json.NewDecoder(r.Body).Decode(&updatedTodo)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call UpdateTodo function
	updatedTodo, err = database.UpdateTodo(id, updatedTodo.Title, updatedTodo.Completed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return updated Todo in response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTodo)
}

// DeleteTodo xóa một công việc
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	err := database.DeleteTodo(id)
	if err != nil {
		http.Error(w, "Không thể xóa công việc", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
