package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *Admin) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u *Admin) CheckPassword(password string) error {
	// komparasi database dengan parameter
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}