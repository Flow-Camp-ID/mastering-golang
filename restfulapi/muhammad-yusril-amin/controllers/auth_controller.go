package controllers

import (
	"restfull-api-lms/models"
	"restfull-api-lms/utils"

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
	var admin models.Admin

	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if  err := admin.HashPassword(admin.Password); err != nil {
		c.JSON(500, gin.H{"Error": "Gagal encrypt password"})
		return
	}

	result := ac.DB.Create(&admin)

	if result.Error != nil {
		c.JSON(400, gin.H{"Error": "Gagal membuat Admin!"})
		return
	}

	token, err := utils.GenerateToken(admin.ID)

	if err != nil {
		c.JSON(400, gin.H{"Error": "Gagal generate token!"})
		return
	}

	c.JSON(201, gin.H{
		"Message": "Admin berhasil dibuat",
		"Token": token,
	})
}

func (ac *AuthController) Login(c *gin.Context) {
	var loginReq models.LoginRequest

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(400, gin.H{"error:": err.Error()})
	}

	var user models.Admin

	if err := ac.DB.Where("email = ?", loginReq.Email).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "Email salah!"})
		return
	}

	// check password function
	if err := user.CheckPassword(loginReq.Password); err != nil {
		c.JSON(401, gin.H{"error": "Password salah!"})
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Gagal generate token!"})
		return
	}

	c.JSON(200, gin.H{
		"Message": "Login successfully!",
		"Token": token,
	})
}