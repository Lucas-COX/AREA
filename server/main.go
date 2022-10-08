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

	// goth.UseProviders(
	// 	google.New(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), "http://localhost:8080/providers/google/callback", "https://www.googleapis.com/auth/gmail.readonly"),
	// )

<<<<<<< HEAD
<<<<<<< HEAD
	// Créer le jobsManager
	// Appeler JobsManager.Run() en go routine
=======
>>>>>>> 7b575bd (feat(server): implement google oauth flow to get refresh token)
=======
	// Créer le jobsManager
	// Appeler JobsManager.Run() en go routine
>>>>>>> d2840a3 (feat(server): move providers into handlers and add jobsManager)
	if os.Getenv("PORT") != "" {
		http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r)
	} else {
		http.ListenAndServe(config.Server.Address, r)
	}
}
