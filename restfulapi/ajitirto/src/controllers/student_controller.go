package controllers

import (
	"fmt"
	"strconv"

	"github.com/ajitirto/restfulapi/src/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StudentController struct {
	DB *gorm.DB
}

type StudentResponse struct {
	ID    uint    `json:"id"`
	Name  *string `json:"name"`
	Email *string `json:"email"`
	Age   *int    `json:"age"`
	Major *string `json:"major"`
}

type UpdateStudentRequest struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Age   int    `json:"age,omitempty"`
	Major string `json:"major,omitempty"`
}

func NewStudentController(db *gorm.DB) *StudentController {
	return &StudentController{DB: db}
}

func (sc *StudentController) CreateStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := sc.DB.Create(&student).Error; err != nil {
		c.JSON(400, gin.H{"error": "Gagal membuat student"})
		return
	}

	response := StudentResponse{
		ID:    student.ID,
		Name:  student.Name,
		Email: student.Email,
		Age:   student.Age,
		Major: student.Major,
	}

	c.JSON(201, gin.H{
		"Message": "student created successfully",
		"Student": response,
	})
}

func (sc *StudentController) GetStudents(c *gin.Context) {
	var students []models.Student

	if err := sc.DB.Find(&students).Error; err != nil {
		c.JSON(400, gin.H{"error": "Gagal menampilkan students"})
		return
	}

	response := make([]StudentResponse, len(students))

	for i, student := range students {
		response[i] = StudentResponse{
			ID:    student.ID,
			Name:  student.Name,
			Email: student.Email,
			Age:   student.Age,
			Major: student.Major,
		}
	}
	c.JSON(200, gin.H{
		"Students": response,
	})
}

func (sc *StudentController) UpdateStudent(c *gin.Context) {
	id := c.Param("studentId")

	var student models.Student

	if err := sc.DB.First(&student, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Student not found"})
		return
	}

	var updateData UpdateStudentRequest

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if updateData.Name != "" {
		student.Name = &updateData.Name
	}
	if updateData.Email != "" {
		student.Email = &updateData.Email
	}
	if updateData.Age != 0 {
		student.Age = &updateData.Age
	}
	if updateData.Major != "" {
		student.Major = &updateData.Major
	}

	if err := sc.DB.Save(&student).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update student"})
		return
	}

	response := StudentResponse{
		ID:    student.ID,
		Name:  student.Name,
		Email: student.Email,
		Age:   student.Age,
		Major: student.Major,
	}

	c.JSON(200, gin.H{
		"Message": "Student updated successfully",
		"Student": response,
	})
}

func (sc *StudentController) DeleteStudent(c *gin.Context) {
	id := c.Param("studentId")

	fmt.Printf("ID yang diterima: '%s'\n", id)
	if id == "" {
		c.JSON(400, gin.H{
			"error": "ID tidak valid",
			"id":    id,
		})
		return
	}

	_, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID harus berupa angka"})
		return
	}

	if err := sc.DB.Where("id = ?", id).Delete(&models.Student{}).Error; err != nil { // Perbaikan di sini
		c.JSON(400, gin.H{"error": "Gagal menghapus student"})
		return
	}

	c.JSON(200, gin.H{
		"Message": "Student deleted successfully",
	})
}
