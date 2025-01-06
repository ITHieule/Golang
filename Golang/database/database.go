package database

import (
	"database/sql"
	"fmt"
	"log"

	"Golang/models"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// InitDB khởi tạo kết nối cơ sở dữ liệu
func InitDB() {
	var err error

	// Thay đổi thông tin kết nối phù hợp với MySQL của bạn
	dsn := "root:123456@tcp(127.0.0.1:3306)/todo_app"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Không thể kết nối đến MySQL: %v", err)
	}

	// Kiểm tra kết nối
	if err = DB.Ping(); err != nil {
		log.Fatalf("Không thể ping MySQL: %v", err)
	}

	fmt.Println("Kết nối MySQL thành công!")
}

// GetTodos trả về danh sách công việc từ cơ sở dữ liệu
func GetTodos() ([]models.Todo, error) {
	rows, err := DB.Query("SELECT id, title, completed FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Completed)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

// CreateTodo thêm một công việc mới vào cơ sở dữ liệu
func CreateTodo(todo models.Todo) (int64, error) {
	result, err := DB.Exec("INSERT INTO todos (title, completed) VALUES (?, ?)", todo.Title, todo.Completed)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func UpdateTodo(id int, title string, completed bool) (models.Todo, error) {
	fmt.Printf("Updating todo: ID=%d, Title=%s, Completed=%v\n", id, title, completed)

	var count int
	query := "SELECT COUNT(*) FROM todos WHERE id = ?"
	err := DB.QueryRow(query, id).Scan(&count)
	if err != nil || count == 0 {
		return models.Todo{}, fmt.Errorf("ID không tồn tại")
	}

	var currentTitle string
	query = "SELECT title FROM todos WHERE id = ?"
	err = DB.QueryRow(query, id).Scan(&currentTitle)
	if err != nil {
		return models.Todo{}, fmt.Errorf("Không thể lấy dữ liệu hiện tại")
	}

	fmt.Printf("Current title: %s\n", currentTitle)

	// Nếu title không thay đổi, chỉ cập nhật completed
	if title == "" || currentTitle == title {
		query = "UPDATE todos SET completed = ? WHERE id = ?"
		_, err = DB.Exec(query, completed, id)
	} else {
		// Nếu title thay đổi, cập nhật cả title và completed
		if title != "" {
			query = "UPDATE todos SET title = ?, completed = ? WHERE id = ?"
			_, err = DB.Exec(query, title, completed, id)
		} else {
			return models.Todo{}, fmt.Errorf("Title không thể rỗng")
		}
	}

	if err != nil {
		fmt.Println("Error updating todo:", err)
		return models.Todo{}, err
	}

	// Trả về Todo đã cập nhật
	var updatedTodo models.Todo
	query = "SELECT id, title, completed FROM todos WHERE id = ?"
	err = DB.QueryRow(query, id).Scan(&updatedTodo.ID, &updatedTodo.Title, &updatedTodo.Completed)
	if err != nil {
		return models.Todo{}, fmt.Errorf("Không thể lấy todo đã cập nhật")
	}

	return updatedTodo, nil
}

// DeleteTodo xóa công việc
func DeleteTodo(id int) error {
	_, err := DB.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}
