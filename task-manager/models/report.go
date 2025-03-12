package models

import "time"

type Report struct {
	ID          uint      `gorm:"primaryKey"`
	ReportType  string    `gorm:"type:enum('pdf', 'csv')"`
	GeneratedAt time.Time `gorm:"autoCreateTime"`
	UserID      uint
}
