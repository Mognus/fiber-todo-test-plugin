package todo

import (
	"time"

	"gorm.io/gorm"
)

// Todo represents a todo item
type Todo struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `gorm:"size:200;not null" json:"title"`
	Completed bool           `gorm:"default:false" json:"completed"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // Soft delete support
}

// TableName - Custom table name
func (Todo) TableName() string {
	return "todos"
}
