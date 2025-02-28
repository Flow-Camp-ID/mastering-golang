package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email"`
	Age string `json:"age"`
	Major string `json:"major"`
	Classes []Class `json:"classes" gorm:"many2many:student_class"`
}