package main

import (
	"Area/config"
	"Area/database"
	"Area/database/models"
	"Area/router"
	"net/http"
)

func main() {
	config := config.Read()
	db := database.New(config)
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Trigger{})
	r := router.New()
	http.ListenAndServe(config.Server.Address, r)
}
