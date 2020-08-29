package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/syukri21/Paperid-Golang-Testcase/src/database/entity"
)

// AutoMigration is auto migrate database
func AutoMigration(conn *gorm.DB) {
	conn.AutoMigrate(entity.User{}, entity.Profile{})
	logrus.Info("Success running migration")
}
