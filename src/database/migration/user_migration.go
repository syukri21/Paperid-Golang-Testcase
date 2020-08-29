package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/syukri21/Paperid-Golang-Testcase/src/database/entity"
)

// CreateUser is create user tabel for migration
func CreateUser(conn *gorm.DB) {
	conn.AutoMigrate(&entity.User{})
	logrus.Info("Success running migration")
}
