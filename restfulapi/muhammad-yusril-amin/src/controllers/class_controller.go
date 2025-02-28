package controllers

import (
	"restfull-api-lms/src/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ClassController struct {
	DB *gorm.DB
}

func NewClassController(db *gorm.DB) *ClassController {
	return &ClassController{DB: db}
}

// get all
func (uc *ClassController) GetClasses(c *gin.Context) {
	var classes []models.Class
	var total int64

	page := c.DefaultQuery("page", "1")
    limit := c.DefaultQuery("limit", "10")

    pageInt, _ := strconv.Atoi(page)
    limitInt, _ := strconv.Atoi(limit)
    offset := (pageInt - 1) * limitInt

    uc.DB.Model(&models.Class{}).Count(&total)

	uc.DB.Limit(limitInt).Offset(offset).Find(&classes)

	c.JSON(200, gin.H{
		"classes": classes,
        "page":  pageInt,
        "limit": limitInt,
        "total": total,
	})
}

// get by id
func (uc *ClassController) GetClassesById(c *gin.Context) {
	id := c.Param("classId") 
	var classes models.Class

	uc.DB.Find(&classes, id)

	c.JSON(200, gin.H{
		"class": classes,
	})
}

// Create
func (rc *ClassController) CreateClass(c *gin.Context) {
	var class models.Class

	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := rc.DB.Create(&class).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"data": class})
}

// Update
func (rc *ClassController) UpdateClass(c *gin.Context) {
	id := c.Param("classId") 
	var class models.Class

	if err := rc.DB.First(&class, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Data Kelas Tidak di Temukan"})
		return
	}

	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := rc.DB.Save(&class).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Edit Data Kelas Berhasil", "data": class})
}

// Delete
func (rc *ClassController) DeleteClass(c *gin.Context) {
	id := c.Param("classId")
	var class models.Class

	if err := rc.DB.First(&class, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Data Kelas Tidak di Temukan"})
		return
	}

	if err := rc.DB.Delete(&class).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Hapus Data Kelas Berhasil"})
}

// add stodent to class
func (rc *ClassController) AddStudentToClass(c *gin.Context) {
	id := c.Param("classId")
	var classStudent struct {
		StudentId uint `json:"student_id"`
	}

	if err := c.ShouldBindJSON(&classStudent); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var class models.Class
	var student models.Student
	
	if err := rc.DB.First(&class, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Data Kelas Tidak Ditemukan"})
		return
	}

	if err := rc.DB.First(&student, classStudent.StudentId).Error; err != nil {
		c.JSON(404, gin.H{"error": "Data Siswa Tidak Ditemukan"})
		return
	}

	if err := rc.DB.Model(&student).Association("Classes").Append(&class); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"Message":"Berhasil menambahkan siswa ke kelas"})
}


// add stodent to class
func (rc *ClassController) RemoveStudentFromClass(c *gin.Context) {
	id := c.Param("classId")
	var classStudent struct {
		StudentId uint `json:"student_id"`
	}

	if err := c.ShouldBindJSON(&classStudent); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var class models.Class
	var student models.Student
	
	if err := rc.DB.First(&class, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Data Kelas Tidak Ditemukan"})
		return
	}

	if err := rc.DB.First(&student, classStudent.StudentId).Error; err != nil {
		c.JSON(404, gin.H{"error": "Data Siswa Tidak Ditemukan"})
		return
	}

	if err := rc.DB.Model(&student).Association("Classes").Delete(&class); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"Message":"Berhasil Menghapus siswa dari kelas"})
}