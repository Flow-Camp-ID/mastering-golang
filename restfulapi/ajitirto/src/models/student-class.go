package models

import "gorm.io/gorm"

type StudentClass struct {
	gorm.Model
	StudentID uint `gorm:"primaryKey" json:"student_id"`
	ClassID   uint `gorm:"primaryKey" json:"class_id"`
	Student   Student `gorm:"foreignKey:StudentID"`
    Class     Class   `gorm:"foreignKey:ClassID"`

}