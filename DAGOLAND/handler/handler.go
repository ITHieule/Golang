package handler

import (
	"DAGOLAND/database"
	"DAGOLAND/model"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	// Decode JSON từ body
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Dữ liệu không hợp lệ", http.StatusBadRequest)
		return
	}

	// Kiểm tra dữ liệu đầu vào
	if user.Username == "" || user.Password == "" || user.Email == "" {
		http.Error(w, "Dữ liệu không đầy đủ", http.StatusBadRequest)
		return
	}

	// Kết nối cơ sở dữ liệu
	database.ConnecDB() // Kết nối CSDL (hàm này không nhận tham số)

	// Thêm người dùng vào cơ sở dữ liệu
	query := "INSERT INTO user (username, email, password) VALUES (?, ?, ?)"
	result, err := database.DB.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		http.Error(w, "Không thể đăng ký người dùng", http.StatusInternalServerError)
		return
	}

	// Lấy ID người dùng vừa tạo
	lastInsertID, _ := result.LastInsertId()
	user.ID = uint(lastInsertID)

	// Trả về phản hồi
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, req *http.Request) {

	var user model.User
	var input model.User
	json.NewDecoder(req.Body).Decode(&input)
	query := "SELECT id, username, email, password FROM user WHERE username = ? AND password = ?"
	err := database.DB.QueryRow(query, input.Username, input.Password).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Failed to login", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
}
func Gettodo(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userID")
	if userID == "" {
		http.Error(w, "Missing user_id", http.StatusBadRequest)
		return
	}
	rows, err := database.DB.Query("SELECT id, user_id, title, description, is_completed, created_at, updated_at FROM todolist WHERE user_id = ?", userID)
	if err != nil {
		http.Error(w, "Failed to retrieve todos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var todos []model.TodoList
	for rows.Next() {
		var todo model.TodoList
		err = rows.Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Description, &todo.IsCompleted, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			http.Error(w, "Failed to parse todos", http.StatusInternalServerError)
			return
		}
		todos = append(todos, todo)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	var todo model.TodoList
	json.NewDecoder(r.Body).Decode(&todo)

	if todo.UserID == 0 || todo.Title == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO todolist (user_id, title, description, is_completed) VALUES (?, ?, ?, ?)"
	result, err := database.DB.Exec(query, todo.UserID, todo.Title, todo.Description, todo.IsCompleted)
	if err != nil {
		http.Error(w, "Failed to add todo", http.StatusInternalServerError)
		return
	}

	lastInsertID, _ := result.LastInsertId()
	todo.ID = uint(lastInsertID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var todo model.TodoList
	json.NewDecoder(r.Body).Decode(&todo)

	query := "UPDATE todolist SET title = ?, description = ?, is_completed = ? WHERE id = ?"
	_, err := database.DB.Exec(query, todo.Title, todo.Description, todo.IsCompleted, id)
	if err != nil {
		http.Error(w, "Failed to update todo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Todo updated"})
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	query := "DELETE FROM todolist WHERE id = ?"
	_, err := database.DB.Exec(query, id)
	if err != nil {
		http.Error(w, "Failed to delete todo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Todo deleted"})
}
