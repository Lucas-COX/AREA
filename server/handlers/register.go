package handlers

import (
	"Area/database"
	"Area/database/models"
	"Area/lib"
	"Area/middleware"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/oauth"
)

func Register(w http.ResponseWriter, r *http.Request) {
	s := oauth.NewBearerServer(
		os.Getenv("TOKEN_SECRET"),
		time.Second*600,
		&middleware.UserVerifier{},
		nil,
	)
	username := r.FormValue("username")
	password := r.FormValue("password")
	r.Form.Add("grant_type", "password")
	err := database.User.Create(models.User{
		Username: username,
		Password: password,
	})
	lib.CheckError(err)
	s.UserCredentials(w, r)
}
