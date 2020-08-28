package main

import (
	"os"

	"github.com/syukri21/Paperid-Golang-Testcase/src/apps"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jpoles1/gopherbadger/logging"
)

func main() {
	r := gin.Default()

	app := new(apps.Application)
	app.CreateApp(r)

	if err := godotenv.Load(); err != nil {
		logging.Error("ENV", err)
	}

	port, found := os.LookupEnv("APP_PORT")
	if !found {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		logging.Error("APP", err)
	}

}
