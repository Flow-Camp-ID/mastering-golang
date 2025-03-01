package models

import "time"

type Class struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"type:varchar(100)"`
	Description string `gorm:"type:text"`
	Teacher     string `gorm:"type:varchar(100)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
