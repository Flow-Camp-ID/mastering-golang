package controllers

import (
	"restfull-api-lms/src/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdminController struct {
	DB *gorm.DB
}

func NewAdminController(db *gorm.DB) *AdminController {
	return &AdminController{DB: db}
}

// get all
func (uc *AdminController) GetAdmins(c *gin.Context) {
	var admins []models.Admin
	var total int64

	page := c.DefaultQuery("page", "1")
    limit := c.DefaultQuery("limit", "10")

    pageInt, _ := strconv.Atoi(page)
    limitInt, _ := strconv.Atoi(limit)
    offset := (pageInt - 1) * limitInt

    // Hitung total data
    uc.DB.Model(&models.Admin{}).Count(&total)

	uc.DB.Limit(limitInt).Offset(offset).Find(&admins)

	c.JSON(200, gin.H{
		"Admin": admins,
        "page":  pageInt,
        "limit": limitInt,
        "total": total,
	})
}

// get by id
func (uc *AdminController) GetAdminsById(c *gin.Context) {
	id := c.Param("id") 
	var admins models.Admin

	uc.DB.Find(&admins, id)

	c.JSON(200, gin.H{
		"Admin": admins,
	})
}

// Create
func (rc *AdminController) CreateAdmin(c *gin.Context) {
	var admin models.Admin

	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if  err := admin.HashPassword(admin.Password); err != nil {
		c.JSON(500, gin.H{"Error": "Gagal encrypt password"})
		return
	}

	if err := rc.DB.Create(&admin).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
		
	}

	c.JSON(201, gin.H{"data": admin})
}

// Update
func (rc *AdminController) UpdateAdmin(c *gin.Context) {
	id := c.Param("id") 
	var admin models.Admin

	if err := rc.DB.First(&admin, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Data Admin Tidak di Temukan"})
		return
	}

	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if  err := admin.HashPassword(admin.Password); err != nil {
		c.JSON(500, gin.H{"Error": "Gagal encrypt password"})
		return
	}

	if err := rc.DB.Save(&admin).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Edit Data Admin Berhasil", "data": admin})
}

// Delete
func (rc *AdminController) DeleteAdmin(c *gin.Context) {
	id := c.Param("id")
	var admin models.Admin

	if err := rc.DB.First(&admin, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Data Admin Tidak di Temukan"})
		return
	}

	if err := rc.DB.Delete(&admin).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Hapus Data Admin Berhasil"})
}