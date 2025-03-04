package controllers

import (
	"resfulapi/models"
	"resfulapi/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {

	return &AuthController{DB: db}

}

func (ac *AuthController) Register(c *gin.Context) {

	var Admin models.Admin

	if err := c.ShouldBindJSON(&Admin); err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}

	// hash password
	if err := Admin.HashPassword(Admin.Password); err != nil {
		c.JSON(500, gin.H{"Error": "Gagal encrypt password"})
		return
	}

	result := ac.DB.Create(&Admin)

	if result.Error != nil {
		c.JSON(400, gin.H{"Error": "Gagal membuat Admin"})
		return
	}

	// admin berhasil dibuat, tidak perlu info token
	c.JSON(201, gin.H{
		"Message": "Admin berhasil dibuat",
	})

}

func (ac *AuthController) Login(c *gin.Context) {

	var loginReq models.LoginRequest

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(400, gin.H{"error:": err.Error()})
	}

	var User models.Admin

	if err := ac.DB.Where("email = ?", loginReq.Email).First(&User).Error; err != nil {
		c.JSON(401, gin.H{"error": "Email salah!"})
		return
	}

	// check password
	if err := User.CheckPassword(loginReq.Password); err != nil {
		c.JSON(401, gin.H{"error": "Password salah!"})
		return
	}

	// generate token
	token, err := utils.GenerateToken(User.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Gagal generate token!"})
		return
	}

	// berhasil login dan menginfokan token
	c.JSON(200, gin.H{
		"Message": "Login berhasil!",
		"Token":   token,
	})

}
