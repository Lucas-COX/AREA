package handlers

import (
	"Area/lib"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	var resp logoutResponseBody
	lib.RemoveCookie(w, "area_token")
	resp.Message = "Successfully logged out"
	lib.SendJson(w, resp)
}
