package controllers

import (
	"github.com/ajitirto/restfulapi/src/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ClassController struct {
	DB *gorm.DB
}

type ClassResponse struct {
	ID         uint     `json:"id"`
	Name       *string `json:"name"`
	Code       *string `json:"code"`
	Instructor *string `json:"instructor"`
	Schedule   *string `json:"schedule"`
}

type UpdateClassRequest struct {
	Name       string `json:"name,omitempty"`
	Code       string `json:"code,omitempty"`
	Instructor string `json:"instructor,omitempty"`
	Schedule   string `json:"schedule,omitempty"`
}

func NewClassController(db *gorm.DB) *ClassController {
	return &ClassController{DB: db}
}

func (cc *ClassController) CreateClass(c *gin.Context) {
	var class models.Class

	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := cc.DB.Create(&class).Error; err != nil {
		c.JSON(400, gin.H{"error": "Gagal membuat class"})
		return
	}

	response := ClassResponse{
		ID:         class.ID,
		Name:       class.Name,
		Code:       class.Code,
		Instructor: class.Instructor,
		Schedule:   class.Schedule,
	}

	c.JSON(201, gin.H{
		"Message": "class created successfully",
		"Class":   response,
	})
}

func (cc *ClassController) GetClasses(c *gin.Context) {
	var classes []models.Class

	if err := cc.DB.Find(&classes).Error; err != nil {
		c.JSON(400, gin.H{"error": "Gagal menampilkan class"})
		return
	}

	response := make([]ClassResponse, len(classes))

	for i, class := range classes {
		response[i] = ClassResponse{
			ID:         class.ID,
			Name:       class.Name,
			Code:       class.Code,
			Instructor: class.Instructor,
			Schedule:   class.Schedule,
		}
	}

	c.JSON(200, gin.H{
		"Classes": response,
	})
}

func (cc *ClassController) DetailClass(c *gin.Context) {
	id := c.Param("classId")

	var class models.Class

	if err := cc.DB.Where("id = ?", id).First(&class).Error; err != nil {
		c.JSON(400, gin.H{"error": "Gagal menampilkan class"})
		return
	}

	response := ClassResponse{
		ID:         class.ID,
		Name:       class.Name,
		Code:       class.Code,
		Instructor: class.Instructor,
		Schedule:   class.Schedule,
	}

	c.JSON(200, gin.H{
		"Class": response,
	})
}

func (cc *ClassController) AddStudentToClass(c *gin.Context) {
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
	
	if err := cc.DB.First(&class, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Data Kelas Tidak Ditemukan"})
		return
	}

	if err := cc.DB.First(&student, classStudent.StudentId).Error; err != nil {
		c.JSON(404, gin.H{"error": "Data Siswa Tidak Ditemukan"})
		return
	}

	if err := cc.DB.Model(&student).Association("Classes").Append(&class); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"Message":"Berhasil menambahkan siswa ke kelas"})
}	

func (cc *ClassController) RemoveStudentFromClass(c *gin.Context) {
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
	
	if err := cc.DB.First(&class, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Data Kelas Tidak Ditemukan"})
		return
	}

	if err := cc.DB.First(&student, classStudent.StudentId).Error; err != nil {
		c.JSON(404, gin.H{"error": "Data Siswa Tidak Ditemukan"})
		return
	}

	if err := cc.DB.Model(&student).Association("Classes").Delete(&class); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"Message":"Berhasil Menghapus siswa dari kelas"})
}

func (cc *ClassController) UpdateClass(c *gin.Context) {
	id := c.Param("classId")	

	var class models.Class

	if err := cc.DB.First(&class, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Data Kelas Tidak Ditemukan"})
		return
	}	

	var updateData UpdateClassRequest

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if updateData.Name != "" {
		class.Name = &updateData.Name
	}
	if updateData.Code != "" {
		class.Code = &updateData.Code
	}
	if updateData.Instructor != "" {
		class.Instructor = &updateData.Instructor
	}
	if updateData.Schedule != "" {
		class.Schedule = &updateData.Schedule
	}

	if err := cc.DB.Save(&class).Error; err != nil {
		c.JSON(400, gin.H{"error": "Gagal mengupdate class"})
		return
	}

	response := ClassResponse{
		ID:         class.ID,
		Name:       class.Name,
		Code:       class.Code,
		Instructor: class.Instructor,
		Schedule:   class.Schedule,
	}

	c.JSON(200, gin.H{
		"Message": "Class updated successfully",
		"Class": response,
	})
}

func (cc *ClassController) DeleteClass(c *gin.Context) {
	id := c.Param("classId")	

	var class models.Class

	if err := cc.DB.First(&class, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Data Kelas Tidak Ditemukan"})
		return
	}	

	if err := cc.DB.Delete(&class).Error; err != nil {
		c.JSON(400, gin.H{"error": "Gagal menghapus class"})
		return
	}

	c.JSON(200, gin.H{
		"Message": "Class deleted successfully",
	})
}