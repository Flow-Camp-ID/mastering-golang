package main

import (
	"lms/src/database"
	"lms/src/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectDB()
	router.SetupRouter(r)
	r.Run(":8080") // Menjalankan server di port 8080
}
