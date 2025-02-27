package utils

import "fmt"

type Logger struct {
	prefix string
}

func (l Logger) Log(message string) {
	fmt.Println("ini log: ", message)
}

type Database struct {
	connection string
}

func (d Database) Query(sql string) string {
	return fmt.Sprintf("Execute %s: ", d.connection)
}

type UserService struct {
	logger Logger
	database Database
}

func NewUserService() UserService {
	return UserService{
		logger: Logger{prefix: "UserService"},
		database: Database{connection: "mysql://localhost"},
	}
}

func (u *UserService) CreateUser(name string) {
	u.logger.Log("Creating user: " + name)
	result := u.database.Query("INSERT INTO users (name) VALUES(" + name + ")")
	u.logger.Log(result)
}

func Composition() {
	UserService := NewUserService() // baris 1
	UserService.CreateUser("Taufik") // baris 2
}	