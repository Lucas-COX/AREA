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
	}
	http.SetCookie(w, cookie)
}
