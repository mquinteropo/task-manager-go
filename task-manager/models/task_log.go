package models

import "time"

type TaskLog struct {
	ID        uint `gorm:"primaryKey"`
	TaskID    uint
	Action    string    `gorm:"type:enum('creado', 'actualizado', 'eliminado')"`
	ChangedAt time.Time `gorm:"autoCreateTime"`
	UserID    uint
}
