package main

import (
	"Area/config"
	"Area/database"
	"Area/database/models"
	"Area/router"
	"fmt"
	"net/http"
	"os"
)

func main() {
	config := config.Read()
	db := database.New(config)
	db.AutoMigrate(&models.User{}, &models.Trigger{}, &models.Action{}, &models.Reaction{})
	r := router.New()
	if os.Getenv("PORT") != "" {
		http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r)
	} else {
		http.ListenAndServe(config.Server.Address, r)
	}
}
