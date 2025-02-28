package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	ID    uint     `gorm:"primary_key" json:"id"`
	Name  *string `json:"name"`
	Email *string `json:"email"`
	Age   *int    `json:"age"`
	Major *string `json:"major"`
}
