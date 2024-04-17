package repositories

import (
	"gorm.io/gorm"
)

type Users struct {
	Id             uint   `json:"id"`
	Email          string `json:"email"`
	Username       string `json:"username"`
	Password       string `json:"-"`
	DepartmentId   int    `json:"department_id"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	ProfileImgPath string `json:"profile_img_path" gorm:"column:profile_img_path"`
	Status         bool   `json:"status"`
	Tel            string `json:"tel"`
	CreatedBy      int    `json:"created_by"`
	UpdatedBy      int    `json:"updated_by"`
}

func GetUserRepo(conn *gorm.DB, data interface{}) (interface{}, error) {
	query := conn
	var user Users

	if err := query.Table("users").Where("status = ?", true).Find(&user).Error; err != nil {
		return "", err
	}
	return user, nil
}

func GetUserByIDRepo(ID int) string {

	return "ศูนย์บาท"
}
