package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword mã hóa mật khẩu
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPassword so sánh mật khẩu đã mã hóa với mật khẩu người dùng nhập
// func CheckPassword(hashedPassword, password string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
// 	return err == nil
// }
