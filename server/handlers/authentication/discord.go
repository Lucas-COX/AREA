package authentication

import (
	"Area/database"
	"Area/lib"
	"net/http"
)

type InfoResponse struct {
	Message string `json:"message"`
}

func DiscordLogin(w http.ResponseWriter, r *http.Request) {
	var res InfoResponse
	var user, err = database.User.GetFromContext(r.Context())

	lib.CheckError(err)

	user.DiscordEnabled = true
	database.User.Update(*user)
	res.Message = "OK"

	lib.SendJson(w, res)
}

func DiscordLogout(w http.ResponseWriter, r *http.Request) {
	var user, err = database.User.GetFromContext(r.Context())

	lib.CheckError(err)

	user.DiscordEnabled = false
	_, err = database.User.Update(*user)
	lib.CheckError(err)
	lib.SendJson(w, InfoResponse{Message: "OK"})
}
