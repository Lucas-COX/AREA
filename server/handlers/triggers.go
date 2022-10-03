package handlers

import (
	"Area/lib"
	"net/http"
)

func Triggers(w http.ResponseWriter, r *http.Request) {
	user, err := UserFromContext(r.Context())
	lib.CheckError(err)
	lib.SendJson(w, user)
}
