package handlers

import (
	"Area/database"
	"Area/database/models"
	"Area/lib"
	"Area/middleware"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/oauth"
	"github.com/go-sql-driver/mysql"
)

func Token(w http.ResponseWriter, r *http.Request) {
	var mysqlErr *mysql.MySQLError
	s := oauth.NewBearerServer(
		os.Getenv("TOKEN_SECRET"),
		time.Second*600,
		&middleware.UserVerifier{},
		nil,
	)
	username := r.FormValue("username")
	password := r.FormValue("password")
	grant_type := r.FormValue("grant_type")
	if grant_type == "register" {
		r.Form.Set("grant_type", "password")
		err := database.User.Create(models.User{
			Username: username,
			Password: password,
		})
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			lib.SendError(w, 400, "User already exists")
			return
		}
		lib.CheckError(err)
	}
	s.UserCredentials(w, r)
}
