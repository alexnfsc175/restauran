package main

import (
	"fmt"
	"log"

	Config "github.com/alexnfsc175/restaurant/config"
	Models "github.com/alexnfsc175/restaurant/models"
	Routes "github.com/alexnfsc175/restaurant/routes"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var err error

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// fmt.Println(os.Getenv("PORT"))
	// fmt.Println(Config.DbURL(Config.BuildDBConfig()))

	Config.DB, err = gorm.Open("postgres", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("statuse: ", err)
	}

	if err = Config.DB.DB().Ping(); err != nil {
		fmt.Println("statuse: ", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Todo{})

	r := Routes.SetupRouter()
	// running
	r.Run()
}
