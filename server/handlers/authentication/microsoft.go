package authentication

import (
	"Area/database"
	"Area/lib"
	"Area/services"
	"encoding/base64"
	"net/http"
)

func MicrosoftLogin(w http.ResponseWriter, r *http.Request) {
	var res urlResponse
	var user, err = database.User.GetFromContext(r.Context())

	callbackUrl, err := base64.RawStdEncoding.DecodeString(r.URL.Query().Get("callback"))
	if err != nil || string(callbackUrl) == "" {
		lib.SendError(w, 400, "Missing callback url (base64 encoded)")
		return
	}

	res.Url = services.Microsoft.Authenticate(string(callbackUrl), user.ID)

	lib.SendJson(w, res)
}

func MicrosoftCallback(w http.ResponseWriter, r *http.Request) {

	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	if code == "" {
		lib.SendError(w, http.StatusBadRequest, "Invalid code provided")
		return
	}
	if state == "" {
		lib.SendError(w, http.StatusBadRequest, "Invalid state provided")
		return
	}

	url, err := services.Microsoft.AuthenticateCallback(state, code)
	if err != nil {
		lib.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Add("Location", url)
	w.WriteHeader(http.StatusPermanentRedirect)
}
