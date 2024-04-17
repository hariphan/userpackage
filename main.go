package user

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func GetUser(conn *gorm.DB, data interface{}) string {
	// repositories.GetUserRepo(conn, data)
	return "ศูนย์บาท"
}

func GetUserByID(ID int) string {

	return "ศูนย์บาท"
}
