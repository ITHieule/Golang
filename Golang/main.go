package main

import (
	"Golang/database"
	"Golang/routes"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	// Kết nối cơ sở dữ liệu
	database.InitDB()
	defer database.DB.Close()

	// Định nghĩa router
	r := routes.SetupRoutes()

	// Cấu hình CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080"}, // Cho phép frontend Vue.js
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	// API endpoint đơn giản mà không có dữ liệu mẫu
	r.HandleFunc("/api/items", func(w http.ResponseWriter, r *http.Request) {
		// Chỉ trả về thông báo JSON mà không có dữ liệu mẫu
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)                     // Trả về mã trạng thái 200 OK
		w.Write([]byte(`{"message": "API is working"}`)) // Trả về thông báo đơn giản
	}).Methods("GET")

	// Sử dụng CORS với router
	handler := c.Handler(r)

	// Chạy server trên cổng 8081
	log.Fatal(http.ListenAndServe(":8081", handler))
}
