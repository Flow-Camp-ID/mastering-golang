package models

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	ID         uint     `gorm:"primary_key" json:"id"`
	Name       *string `json:"name"`
	Code       *string `json:"code"`
	Instructor *string `json:"instructor"`
	Schedule   *string `json:"schedule"`
}
