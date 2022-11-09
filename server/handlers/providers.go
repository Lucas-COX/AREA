package handlers

import (
	"Area/handlers/authentication"
	"Area/lib"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ProviderLogin(w http.ResponseWriter, r *http.Request) {
	provider, err := authentication.Parse(chi.URLParam(r, "provider"))

	if err != nil {
		lib.SendError(w, http.StatusBadRequest, "Invalid provider.")
	}

	switch provider {
	case authentication.Google:
		authentication.GoogleLogin(w, r)
	case authentication.Microsoft:
		authentication.MicrosoftLogin(w, r)
	case authentication.Github:
		authentication.GithubLogin(w, r)
	case authentication.Notion:
		authentication.NotionLogin(w, r)
	case authentication.Discord:
		authentication.DiscordLogin(w, r)
	case authentication.Timer:
		authentication.TimerLogin(w, r)
	default:
		lib.SendError(w, http.StatusBadRequest, "Invalid provider.")
	}
}

func ProviderCallback(w http.ResponseWriter, r *http.Request) {
	provider, err := authentication.Parse(chi.URLParam(r, "provider"))
	if err != nil {
		lib.SendError(w, http.StatusBadRequest, "Invalid provider.")
	}

	switch provider {
	case authentication.Google:
		authentication.GoogleCallback(w, r)
	case authentication.Microsoft:
		authentication.MicrosoftCallback(w, r)
	case authentication.Github:
		authentication.GithubCallback(w, r)
	case authentication.Notion:
		authentication.NotionCallback(w, r)
	default:
		lib.SendError(w, http.StatusBadRequest, "Invalid profiler")
	}
}

func ProviderLogout(w http.ResponseWriter, r *http.Request) {
	provider, err := authentication.Parse(chi.URLParam(r, "provider"))

	if err != nil {
		lib.SendError(w, http.StatusBadRequest, "Invalid provider.")
	}

	switch provider {
	case authentication.Google:
		authentication.GoogleLogout(w, r)
	case authentication.Microsoft:
		authentication.MicrosoftLogout(w, r)
	case authentication.Github:
		authentication.GithubLogout(w, r)
	case authentication.Notion:
		authentication.NotionLogout(w, r)
	case authentication.Discord:
		authentication.DiscordLogout(w, r)
	case authentication.Timer:
		authentication.TimerLogout(w, r)
	}
}
