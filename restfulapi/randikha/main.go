package main

import (
	"log"
	"resfulapi/config"
	"resfulapi/controllers"
	"resfulapi/middleware"
	"resfulapi/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("ENV gagal terbaca")
	}

	//gin init
	r := gin.Default()

	// database connect
	db := config.ConnectDatabase()

	// data migrate
	db.AutoMigrate(&models.Admin{})
	db.AutoMigrate(&models.Student{})
	db.AutoMigrate(&models.Class{})

	// initial controller
	AuthController := controllers.NewAuthController(db)
	AdminController := controllers.NewAdminController(db)
	StudentController := controllers.NewStudentController(db)
	ClassController := controllers.NewClassController(db)

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("register", AuthController.Register)
			auth.POST("login", AuthController.Login)
		}

		protected := api.Group("/")
		protected.Use(middleware.AuthMidlleware())
		{
			admin := protected.Group("/admin")
			{
				admin.GET("/", AdminController.GetAdmins)
				admin.GET("/:adminId", AdminController.GetAdminsById)
				admin.PUT("/:adminId", AdminController.UpdateAdmin)
				admin.DELETE("/:adminId", AdminController.DeleteAdmin)
			}

			student := protected.Group("/student")
			{
				student.GET("/", StudentController.GetStudents)
				student.GET("/:studentId", StudentController.GetStudentById)
				student.POST("/", StudentController.CreateStudent)
				student.PATCH("/:studentId", StudentController.UpdateStudent)
				student.DELETE("/:studentId", StudentController.DeleteStudent)
			}

			class := protected.Group("/class")
			{
				class.GET("/", ClassController.GetClasses)
				class.GET("/:classId", ClassController.GetClassesById)
				class.POST("/", ClassController.CreateClass)
				class.PUT("/:classId", ClassController.UpdateClass)
				class.DELETE("/:classId", ClassController.DeleteClass)
				class.POST("/:classId/students", ClassController.AddStudentToClass)
				class.DELETE("/:classId/students", ClassController.RemoveStudentFromClass)
			}

		}
	}

	r.Run(":3000")
}
