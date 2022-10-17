package main

import (
	"Area/config"
	"Area/database"
	"Area/database/models"
	"Area/jobs"
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

	// Uncomment to create the actions for the first time
	// db.Create(&models.Action{
	// 	Type:  "gmail",
	// 	Event: "receive",
	// })
	// db.Create(&models.Reaction{
	// 	Type:   "discord",
	// 	Action: "send",
	// })

	manager := jobs.NewManager()
	manager.RunAsync()
	if os.Getenv("PORT") != "" {
		http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r)
	} else {
		http.ListenAndServe(config.Server.Address, r)
	}
}
