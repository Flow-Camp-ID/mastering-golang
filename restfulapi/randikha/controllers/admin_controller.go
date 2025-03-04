package controllers

import (
	"resfulapi/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdminController struct {
	DB *gorm.DB
}

func NewAdminController(db *gorm.DB) *AdminController {
	return &AdminController{DB: db}
}

// Get list
func (uc *AdminController) GetAdmins(c *gin.Context) {

	var Admins []models.Admin
	uc.DB.Find(&Admins)

	c.JSON(200, gin.H{
		"Admin": Admins,
	})
}

// Get By Id
func (uc *AdminController) GetAdminsById(c *gin.Context) {
	id := c.Param("id")
	var Admins models.Admin

	uc.DB.Find(&Admins, id)

	c.JSON(200, gin.H{
		"Admin": Admins,
	})
}

// Update
func (rc *AdminController) UpdateAdmin(c *gin.Context) {

	id := c.Param("id")
	var Admin models.Admin

	if err := rc.DB.First(&Admin, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&Admin); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := Admin.HashPassword(Admin.Password); err != nil {
		c.JSON(500, gin.H{"Error": "Gagal encrypt password"})
		return
	}

	if err := rc.DB.Save(&Admin).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Update data berhasil", "data": Admin})

}

// Delete
func (rc *AdminController) DeleteAdmin(c *gin.Context) {

	id := c.Param("id")
	var Admin models.Admin

	if err := rc.DB.First(&Admin, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	if err := rc.DB.Delete(&Admin).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Detele data berhasil"})

}
