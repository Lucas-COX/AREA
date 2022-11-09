package authentication

import (
	"Area/database"
	"Area/lib"
	"net/http"
)

func TimerLogin(w http.ResponseWriter, r *http.Request) {
	var res InfoResponse
	var user, err = database.User.GetFromContext(r.Context())

	lib.CheckError(err)

	user.TimerEnabled = true
	database.User.Update(*user)
	res.Message = "OK"

	lib.SendJson(w, res)
}

func TimerLogout(w http.ResponseWriter, r *http.Request) {
	var user, err = database.User.GetFromContext(r.Context())

	lib.CheckError(err)

	user.TimerEnabled = false
	_, err = database.User.Update(*user)
	lib.CheckError(err)
	lib.SendJson(w, InfoResponse{Message: "OK"})
}
