package main

import (
	"fmt"
	"os"

	"github.com/evaldasNe/stock-portfolio-web/Config"
	"github.com/evaldasNe/stock-portfolio-web/Models"
	"github.com/evaldasNe/stock-portfolio-web/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})
	r := Routes.SetupRouter()
	//running
	r.Run(":" + os.Getenv("APP_PORT"))
}
