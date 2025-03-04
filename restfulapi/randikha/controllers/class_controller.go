package controllers

import (
	"resfulapi/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ClassController struct {
	DB *gorm.DB
}

func NewClassController(db *gorm.DB) *ClassController {

	return &ClassController{DB: db}

}

// Get List
func (uc *ClassController) GetClasses(c *gin.Context) {

	var Classes []models.Class
	uc.DB.Find(&Classes)

	c.JSON(200, gin.H{
		"Classes": Classes,
	})

}

// Get By Id
func (uc *ClassController) GetClassesById(c *gin.Context) {

	id := c.Param("ClassId")
	var Classes models.Class

	uc.DB.Find(&Classes, id)

	c.JSON(200, gin.H{
		"Classes": Classes,
	})

}

// Create
func (rc *ClassController) CreateClass(c *gin.Context) {

	var Class models.Class

	if err := c.ShouldBindJSON(&Class); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := rc.DB.Create(&Class).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"data": Class})

}

// Update
func (rc *ClassController) UpdateClass(c *gin.Context) {

	id := c.Param("ClassId")
	var Class models.Class

	if err := rc.DB.First(&Class, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&Class); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := rc.DB.Save(&Class).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Update data berhasil", "data": Class})

}

// Delete
func (rc *ClassController) DeleteClass(c *gin.Context) {

	id := c.Param("ClassId")
	var Class models.Class

	if err := rc.DB.First(&Class, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	if err := rc.DB.Delete(&Class).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Delete data berhasil"})

}
