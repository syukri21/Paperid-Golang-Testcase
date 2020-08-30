package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/syukri21/Paperid-Golang-Testcase/src/database/entity"
)

// AutoMigration is auto migrate database
func AutoMigration(conn *gorm.DB) {
	conn.AutoMigrate(entity.User{})

	if conn.HasTable(&entity.User{}) {
		conn.AutoMigrate(entity.Profile{})
		conn.Model(&entity.Profile{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

		conn.AutoMigrate(entity.FinanceAccount{})
		conn.Model(&entity.FinanceAccount{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

		if conn.HasTable(&entity.FinanceAccount{}) {
			conn.AutoMigrate(entity.FinanceTransaction{})
			conn.Model(&entity.FinanceTransaction{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
			conn.Model(&entity.FinanceTransaction{}).AddForeignKey("finance_account_id", "finance_accounts(id)", "CASCADE", "CASCADE")
		}
	}

	conn.AutoMigrate(entity.FinanceAccountType{})
	if (conn.HasTable(&entity.FinanceAccount{})) {
		conn.Model(&entity.FinanceAccount{}).AddForeignKey("type_id", "finance_account_types(id)", "CASCADE", "CASCADE")
	}

	logrus.Info("Success running migration")
}
