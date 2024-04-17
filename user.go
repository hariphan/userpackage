package user

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/hariphan/userpackage/repositories"
	"gorm.io/gorm"
)

func GetUser(conn *gorm.DB, data interface{}) (interface{}, error) {
	res, err := repositories.GetUserRepo(conn, data)
	// return "ศูนย์บาท"
	return res, err
}

func GetUserByID(ID int) string {

	return "ศูนย์บาท"
}
