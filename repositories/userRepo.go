package repositories

import (
	"gorm.io/gorm"
)

func GetUserRepo(conn *gorm.DB, data interface{}) (interface{}, error) {
	query := conn
	if err := query.Where("status = ?", true).Find(&data).Error; err != nil {
		return "", err
	}
	return data, nil
}

func GetUserByIDRepo(ID int) string {

	return "ศูนย์บาท"
}
