package router

import (
	"lms/src/controllers"
	"lms/src/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.POST("/login", controllers.Login)

	// Routes yang memerlukan autentikasi
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/students", controllers.GetStudents)
		auth.POST("/students", controllers.CreateStudent)
		auth.PUT("/students/:id", controllers.UpdateStudent)
		auth.DELETE("/students/:id", controllers.DeleteStudent)

		auth.GET("/classes", controllers.GetClasses)
		auth.POST("/classes", controllers.CreateClass)
		auth.PUT("/classes/:id", controllers.UpdateClass)
		auth.DELETE("/classes/:id", controllers.DeleteClass)
	}
}
