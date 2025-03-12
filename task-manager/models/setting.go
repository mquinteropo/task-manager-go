package models

type Setting struct {
	ID          uint   `gorm:"primaryKey"`
	ConfigKey   string `gorm:"type:varchar(255);unique;not null"`
	ConfigValue string `gorm:"type:text;not null"`
}
