package controllers

import (
	"lms/src/database"
	"lms/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(http.StatusOK, gin.H{"students": students})
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusOK, gin.H{"student": student})
}

func UpdateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&student)
	c.JSON(http.StatusOK, gin.H{"student": student})
}

func DeleteStudent(c *gin.Context) {
	var student models.Student
	if err := database.DB.Where("id = ?", c.Param("id")).First(&student).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found!"})
		return
	}
	database.DB.Delete(&student)
	c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully!"})
}
