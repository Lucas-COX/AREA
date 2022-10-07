package handlers

import (
<<<<<<< HEAD
	"Area/handlers/providers"
	"Area/lib"
=======
	"Area/lib"
	"Area/providers"
>>>>>>> 7b575bd (feat(server): implement google oauth flow to get refresh token)
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ProviderLogin(w http.ResponseWriter, r *http.Request) {
	provider, err := providers.Parse(chi.URLParam(r, "provider"))
	if err != nil {
		lib.SendError(w, http.StatusBadRequest, "Invalid provider.")
	}

	switch provider {
	case providers.Google:
		providers.GoogleLogin(w, r)
		// break
	default:
		lib.SendError(w, http.StatusBadRequest, "Invalid provider.")
	}
}

func ProviderCallback(w http.ResponseWriter, r *http.Request) {
	provider, err := providers.Parse(chi.URLParam(r, "provider"))
	if err != nil {
		lib.SendError(w, http.StatusBadRequest, "Invalid provider.")
	}

	switch provider {
	case providers.Google:
		providers.GoogleCallback(w, r)
		// break
	default:
		lib.SendError(w, http.StatusBadRequest, "Invalid profiler")
	}
}
