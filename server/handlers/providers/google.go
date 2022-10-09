package providers

import (
	"Area/database"
	"Area/lib"
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type urlResponse struct {
	Url string `json:"url"`
}

type oauthState struct {
	Callback string `json:"callback"`
	UserId   uint   `json:"user_id"`
}

func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	var res urlResponse
	var state oauthState
	var user, err = database.User.GetFromContext(r.Context())

	callbackUrl, err := base64.RawStdEncoding.DecodeString(r.URL.Query().Get("callback"))
	if err != nil || string(callbackUrl) == "" {
		lib.SendError(w, 400, "Missing callback url (base64 encoded)")
		return
	}
	state.Callback = r.URL.Query().Get("callback")
	state.UserId = user.ID

	var conf = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8080/providers/google/callback", // add source query parameter
		Scopes: []string{
			"https://www.googleapis.com/auth/gmail.readonly",
		},
		Endpoint: google.Endpoint,
	}

	bytes, _ := json.Marshal(state)
	str := base64.RawStdEncoding.EncodeToString(bytes)
	res.Url = conf.AuthCodeURL(str, oauth2.AccessTypeOffline)
	lib.SendJson(w, res)
}

func GoogleCallback(w http.ResponseWriter, r *http.Request) {
	var state oauthState

	bytes, err := base64.RawStdEncoding.DecodeString(r.URL.Query().Get("state"))
	err2 := json.Unmarshal(bytes, &state)

	if _, err3 := url.Parse(string(state.Callback)); err != nil || err2 != nil || err3 != nil {
		lib.SendError(w, http.StatusBadRequest, "Invalid callback url")
		return
	}

	bytes, err = base64.RawStdEncoding.DecodeString(string(state.Callback))
	if err != nil {
		lib.SendError(w, http.StatusBadRequest, "Invalid callback url")
		return
	}
	state.Callback = string(bytes)

	user, err := database.User.GetById(state.UserId, false)
	if err != nil {
		lib.SendError(w, http.StatusBadRequest, "Invalid user id")
		return
	}

	var conf = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8080/providers/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/gmail.readonly",
		},
		Endpoint: google.Endpoint,
	}

	code := r.URL.Query().Get("code")
	token, err := conf.Exchange(context.Background(), code)

	if err != nil {
		lib.SendError(w, http.StatusBadRequest, "Invalid code provided")
		return
	}

	// Todo: handle "error" query parameter
	// Todo: refactor this process

	// Todo: crypt that token
	user.GoogleToken = token.RefreshToken
	database.User.Update(*user)

	// Client pour faire des appels à l'api google
	// client := conf.Client(context.Background(), token)

	w.Header().Add("Location", string(state.Callback))
	w.WriteHeader(http.StatusPermanentRedirect)
}
