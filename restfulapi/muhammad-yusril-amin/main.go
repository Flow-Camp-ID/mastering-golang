package main

import (
	"log"
	"restfull-api-lms/config"
	"restfull-api-lms/controllers"
	"restfull-api-lms/middleware"
	"restfull-api-lms/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("ENV gagal Terbaca")
	}

	// gin init
	r := gin.Default()

	// database connect
	db := config.ConnectDatabase()
	
	// auto migrate
	db.AutoMigrate(&models.Admin{}, &models.Student{}, &models.Class{})

	// inisialisasi controllers
	authController := controllers.NewAuthController(db)
	adminController := controllers.NewAdminController(db)
	studentController := controllers.NewStudentController(db)
	classController := controllers.NewClassController(db)

	api := r.Group("/api") 
	{
		auth := api.Group("/admin")
		{
			auth.POST("register", authController.Register)
			auth.POST("login", authController.Login)
		}

		protected := api.Group("/")
		protected.Use(middleware.AuthMidlleware()) 
		{
			admin := protected.Group("/admin")
			{
				admin.GET("/", adminController.GetAdmins)
				admin.GET("/:id", adminController.GetAdminsById)
				admin.POST("/store", adminController.CreateAdmin)
				admin.PUT("/:id/update", adminController.UpdateAdmin)
				admin.DELETE("/destroy/:id", adminController.DeleteAdmin)
			}
			
			student := protected.Group("/student")
			{
				student.GET("/", studentController.GetStudents)
				student.GET("/:studentId", studentController.GetStudentById)
				student.POST("/", studentController.CreateStudent)
				student.PATCH("/:studentId", studentController.UpdateStudent)
				student.DELETE("/:studentId", studentController.DeleteStudent)
			}

			class := protected.Group("/class")
			{
				class.GET("/", classController.GetClasses)
				class.GET("/:classId", classController.GetClassesById)
				class.POST("/", classController.CreateClass)
				class.PUT("/:classId", classController.UpdateClass)
				class.DELETE("/:classId", classController.DeleteClass)
				class.POST("/:classId/students", classController.AddStudentToClass)
				class.DELETE("/:classId/students", classController.RemoveStudentFromClass)
			}

		}
	}

	r.Run(":8000")
}