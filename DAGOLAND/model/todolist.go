package model

type TodoList struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null"` // Khóa ngoại
	Title       string `gorm:"not null"`
	Description string
	IsCompleted bool   `gorm:"default:false"`
	CreatedAt   string `gorm:"autoCreateTime"`
	UpdatedAt   string `gorm:"autoUpdateTime"`
}
