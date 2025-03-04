package controllers

import (
	"resfulapi/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StudentController struct {
	DB *gorm.DB
}

func NewStudentController(db *gorm.DB) *StudentController {

	return &StudentController{DB: db}

}

// Get List
func (uc *StudentController) GetStudents(c *gin.Context) {

	var Students []models.Student
	uc.DB.Find(&Students)

	c.JSON(200, gin.H{
		"Students": Students,
	})

}

// Get By Id
func (uc *StudentController) GetStudentById(c *gin.Context) {

	id := c.Param("studentId")
	var Students models.Student

	uc.DB.Find(&Students, id)

	c.JSON(200, gin.H{
		"Students": Students,
	})

}

// Create
func (rc *StudentController) CreateStudent(c *gin.Context) {

	var Student models.Student

	if err := c.ShouldBindJSON(&Student); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := rc.DB.Create(&Student).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"data": Student})

}

// Update
func (rc *StudentController) UpdateStudent(c *gin.Context) {

	id := c.Param("studentId")
	var Student models.Student

	if err := rc.DB.First(&Student, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&Student); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := rc.DB.Save(&Student).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Update data berhasil", "data": Student})

}

// Delete
func (rc *StudentController) DeleteStudent(c *gin.Context) {

	id := c.Param("studentId")
	var Student models.Student

	if err := rc.DB.First(&Student, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	if err := rc.DB.Delete(&Student).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Delete data berhasil"})

}

// add stodent to class
func (rc *ClassController) AddStudentToClass(c *gin.Context) {

	id := c.Param("classId")
	var classStudent struct {
		StudentId uint `json:"studentId"`
	}

	if err := c.ShouldBindJSON(&classStudent); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var class models.Class
	var student models.Student

	if err := rc.DB.First(&class, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Data kelas tidak ditemukan"})
		return
	}

	if err := rc.DB.First(&student, classStudent.StudentId).Error; err != nil {
		c.JSON(404, gin.H{"error": "Data siswa tidak Ditemukan"})
		return
	}

	if err := rc.DB.Model(&student).Association("Classes").Append(&class); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"Message": "Berhasil menambahkan siswa ke kelas"})
}

// Remove stodent from class
func (rc *ClassController) RemoveStudentFromClass(c *gin.Context) {

	id := c.Param("classId")
	var classStudent struct {
		StudentId uint `json:"studentId"`
	}

	if err := c.ShouldBindJSON(&classStudent); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var class models.Class
	var student models.Student

	if err := rc.DB.First(&class, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Data kelas tidak ditemukan"})
		return
	}

	if err := rc.DB.First(&student, classStudent.StudentId).Error; err != nil {
		c.JSON(404, gin.H{"error": "Data siswa tidak Ditemukan"})
		return
	}

	if err := rc.DB.Model(&student).Association("Classes").Delete(&class); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"Message": "Berhasil menghapus siswa dari kelas"})
}
