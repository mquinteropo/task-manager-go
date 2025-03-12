package models

import (
	"time"
)

type Task struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text"`
	Status      string `gorm:"type:enum('pendiente', 'en progreso', 'completada');default:'pendiente'"`
	DueDate     time.Time
	UserID      uint
	User        User `gorm:"foreignKey:UserID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
