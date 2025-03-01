package controllers

import (
	"lms/src/database"
	"lms/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetClasses(c *gin.Context) {
	var classes []models.Class
	database.DB.Find(&classes)
	c.JSON(http.StatusOK, gin.H{"classes": classes})
}

func CreateClass(c *gin.Context) {
	var class models.Class
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&class)
	c.JSON(http.StatusOK, gin.H{"class": class})
}

func UpdateClass(c *gin.Context) {
	var class models.Class
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&class)
	c.JSON(http.StatusOK, gin.H{"class": class})
}

func DeleteClass(c *gin.Context) {
	var class models.Class
	if err := database.DB.Where("id = ?", c.Param("id")).First(&class).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Class not found!"})
		return
	}
	database.DB.Delete(&class)
	c.JSON(http.StatusOK, gin.H{"message": "Class deleted successfully!"})
}
