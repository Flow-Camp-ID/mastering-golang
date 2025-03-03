package controllers

import (
	"github.com/ajitirto/restfulapi/src/models"
	"github.com/ajitirto/restfulapi/src/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

type AdminResponse struct {
    ID    uint   `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{DB: db}
}

func (ac *AuthController) Register(c *gin.Context) {

	var admin models.Admin

	if ac.DB == nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := admin.HashPassword(admin.Password); err != nil {
		c.JSON(500, gin.H{"Error": "Gagal encrypt password!"})
		return
	}

	result := ac.DB.Create(&admin)

	if result.Error != nil {
		c.JSON(400, gin.H{"Error": "Gagal membuat admin users!"})
		return
	}

	_, err := utils.GenerateToken(admin.ID)

	if err != nil {
		c.JSON(400, gin.H{"Error": "Gagal generate token!"})
		return
	}

	response := AdminResponse{
        ID:    admin.ID,
        Name:  admin.Name,
        Email: admin.Email,
    }

	c.JSON(201, gin.H{
		"Message": "Admin registered successfully",
		"Admin":   response,
	})
}

func (ac *AuthController) Login(c *gin.Context) {
	var loginReq models.LoginRequest

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(400, gin.H{"error:": err.Error()})
	}

	var user models.Admin

	if err := ac.DB.Where("email = ?", loginReq.Email).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"Error": "Email salah!"})
		return
	}

	if err := user.CheckPassword(loginReq.Password); err != nil {
		c.JSON(401, gin.H{"Error": "Password salah!"})
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(500, gin.H{"Error": "Gagal generate token!"})
		return
	}

	c.JSON(200, gin.H{
		"Message": "Login successful",
		"Token":   token,
	})
}
