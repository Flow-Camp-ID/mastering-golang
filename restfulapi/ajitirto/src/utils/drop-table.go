package utils

import (
	"github.com/ajitirto/restfulapi/src/config"
	"github.com/ajitirto/restfulapi/src/models"
)

func DropTable() {
	db := config.ConnectMysql()
	db.Migrator().DropTable(&models.Admin{}, &models.Class{}, &models.Student{}, &models.StudentClass{})
}
