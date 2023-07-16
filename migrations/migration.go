package main

import (
	"fmt"
	"log"

	"github.com/Atoo35/basic-crud/configurations"
	"github.com/Atoo35/basic-crud/models"
)

func init() {
	config, err := configurations.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load variables from .env file")
	}
	configurations.Connect(&config)
}

func main() {
	configurations.DB.AutoMigrate(&models.User{}, &models.Book{})
	fmt.Println("Migration completed")
}
