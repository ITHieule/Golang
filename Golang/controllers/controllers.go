package controllers

import (
	"Golang/database"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your_secret_key")

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Dữ liệu không hợp lệ", http.StatusBadRequest)
		return
	}

	var hashedPassword string
	var id int
	query := "SELECT id, password FROM account WHERE username = ?"
	err = database.DB.QueryRow(query, req.Username).Scan(&id, &hashedPassword)
	if err == sql.ErrNoRows {
		http.Error(w, "Tài khoản không tồn tại", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Lỗi hệ thống", http.StatusInternalServerError)
		return
	}

	// if !utils.CheckPassword(hashedPassword, req.Password) {
	// 	http.Error(w, "Sai mật khẩu", http.StatusUnauthorized)
	// 	return
	// }

	token, err := generateJWT(id)
	if err != nil {
		http.Error(w, "Không thể tạo token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LoginResponse{Token: token})
}

func generateJWT(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
