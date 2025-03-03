package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/ajitirto/restfulapi/src/config"
	"github.com/ajitirto/restfulapi/src/controllers"
	"github.com/ajitirto/restfulapi/src/middleware"
	"github.com/ajitirto/restfulapi/src/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error load .env file")
	}

	utils.ConnectToMysql()
	// utils.DropTable()

	router := gin.New()

	// Logger custom
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		colorReset := "\033[0m"
		colorGreen := "\033[32m"  // Hijau untuk 2xx
		colorYellow := "\033[33m" // Kuning untuk 3xx
		colorRed := "\033[31m"    // Merah untuk 4xx dan 5xx

		// Tentukan warna berdasarkan status kode
		statusColor := colorReset
		switch {
		case param.StatusCode >= 200 && param.StatusCode < 300:
			statusColor = colorGreen
		case param.StatusCode >= 300 && param.StatusCode < 400:
			statusColor = colorYellow
		case param.StatusCode >= 400:
			statusColor = colorRed
		}
		return fmt.Sprintf("%s - [%s] \"%s %s %s %s%d%s %s \"%s\" %s\"\n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				statusColor, param.StatusCode, colorReset,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
		)
	}))

	// Logging to a file.
    f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	authController := controllers.NewAuthController(config.ConnectMysql())
	studentController := controllers.NewStudentController(config.ConnectMysql())
	classController := controllers.NewClassController(config.ConnectMysql())

	api := router.Group("/api")
	{
		admin := api.Group("/admin")
		{
			admin.POST("/register", authController.Register )
			admin.POST("/login", authController.Login)
		}

		student := api.Group("/student")
		student.Use(middleware.AuthMiddleware())
		{
			student.GET("", studentController.GetStudents)
			student.POST("", studentController.CreateStudent)
			student.PATCH("/:studentId", studentController.UpdateStudent)
			student.DELETE("/:studentId", studentController.DeleteStudent)
		}

		class := api.Group("/class")
		class.Use(middleware.AuthMiddleware())
		{
			class.POST("", classController.CreateClass)
			class.GET("", classController.GetClasses)
			class.GET("/:classId", classController.DetailClass)
			class.POST("/:classId/students", classController.AddStudentToClass)
			class.DELETE("/:classId/students", classController.RemoveStudentFromClass)
			class.PUT("/:classId", classController.UpdateClass)
			class.DELETE("/:classId", classController.DeleteClass)
		}
	}

	router.Run(":8080")

}
