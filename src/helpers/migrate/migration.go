package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	db "github.com/syukri21/Paperid-Golang-Testcase/src/database"
	"github.com/syukri21/Paperid-Golang-Testcase/src/database/migration"
)

func main() {
	db.AppConnection()
	conn := db.GetDB()
	defer conn.Close()

	migration.AutoMigration(conn)
}
