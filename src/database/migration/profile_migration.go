package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/syukri21/Paperid-Golang-Testcase/src/database/entity"
)

// CreateProfile is create profile tabel for migration
func CreateProfile(conn *gorm.DB) {
	conn.AutoMigrate(&entity.Profile{})
	logrus.Info("Success running migration")
}
