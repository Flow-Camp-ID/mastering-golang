package controllers

import (
	"restfull-api-lms/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StudentController struct {
	DB *gorm.DB
}

func NewStudentController(db *gorm.DB) *StudentController {
	return &StudentController{DB: db}
}


// get all
func (uc *StudentController) GetStudents(c *gin.Context) {
	var students []models.Student
	var total int64

	page := c.DefaultQuery("page", "1")
    limit := c.DefaultQuery("limit", "10")

    pageInt, _ := strconv.Atoi(page)
    limitInt, _ := strconv.Atoi(limit)
    offset := (pageInt - 1) * limitInt

    uc.DB.Model(&models.Student{}).Count(&total)

	uc.DB.Limit(limitInt).Offset(offset).Find(&students)

	c.JSON(200, gin.H{
		"Students": students,
        "page":  pageInt,
        "limit": limitInt,
        "total": total,
	})
}

// get by id
func (uc *StudentController) GetStudentById(c *gin.Context) {
	id := c.Param("studentId") 
	var students models.Student

	uc.DB.Find(&students, id)

	c.JSON(200, gin.H{
		"Students": students,
	})
}

// Create
func (rc *StudentController) CreateStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := rc.DB.Create(&student).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"data": student})
}

// Update
func (rc *StudentController) UpdateStudent(c *gin.Context) {
	id := c.Param("studentId") 
	var student models.Student

	if err := rc.DB.First(&student, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Data Siswa Tidak di Temukan"})
		return
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := rc.DB.Save(&student).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Edit Data Siswa Berhasil", "data": student})
}

// Delete
func (rc *StudentController) DeleteStudent(c *gin.Context) {
	id := c.Param("studentId")
	var student models.Student

	if err := rc.DB.First(&student, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Data Siswa Tidak di Temukan"})
		return
	}

	if err := rc.DB.Delete(&student).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Hapus Data Siswa Berhasil"})
}