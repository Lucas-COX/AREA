package main

import (
	"Area/config"
	"Area/database"
	"Area/database/models"
)

func main() {
	config := config.Read()
	db := database.New(config)
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Trigger{})
}
