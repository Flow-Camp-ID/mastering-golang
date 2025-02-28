package utils

import (
	"github.com/ajitirto/restfulapi/src/config"
	"github.com/ajitirto/restfulapi/src/models"
)

func ConnectToMysql() {
	db := config.ConnectMysql()
	db.AutoMigrate(&models.Admin{}, &models.Class{}, &models.Student{}, &models.StudentClass{})
}