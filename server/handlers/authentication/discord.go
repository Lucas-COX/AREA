package authentication

import (
	"Area/database"
	"Area/lib"
	"net/http"
)

type discordLoginResponse struct {
	Message string `json:"message"`
}

func DiscordLogin(w http.ResponseWriter, r *http.Request) {
	var res discordLoginResponse
	var user, err = database.User.GetFromContext(r.Context())

	lib.CheckError(err)

	user.DiscordEnabled = true
	database.User.Update(*user)
	res.Message = "OK"

	lib.SendJson(w, res)
}
