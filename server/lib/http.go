package lib

import (
	"encoding/json"
	"net/http"
)

func SendJson(w http.ResponseWriter, body any) {
	bytes, err := json.Marshal(body)
	CheckError(err)
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func SetCookie(w http.ResponseWriter, name string, value string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   0,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	}
	http.SetCookie(w, cookie)
}

func RemoveCookie(w http.ResponseWriter, name string) {
	// remove the "area_token" cookie by setting its MaxAge to -1
	cookie := &http.Cookie{
		Name:     name,
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	}
	http.SetCookie(w, cookie)
}
