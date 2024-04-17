package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_get_user(t *testing.T) {
	expected := "ศูนย์บาท"
	// baht := 0.0
	var conn *gorm.DB
	var data interface{}
	actual := GetUser(conn, data)

	assert.Equal(t, expected, actual)
}

func Test_get_user_by_id(t *testing.T) {
	expected := "ศูนย์บาท"
	// baht := 0.0

	actual := GetUserByID(1)

	assert.Equal(t, expected, actual)
}
