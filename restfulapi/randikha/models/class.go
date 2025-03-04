package models

import (
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name       string    `json:"name"`
	Code       string    `json:"code"`
	Instructor string    `json:"instructor"`
	Schedule   string    `json:"schedule"`
	Students   []Student `json:"students" gorm:"many2many:student_class"`
}
