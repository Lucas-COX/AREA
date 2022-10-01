package handlers

import (
	"Area/lib"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	lib.SendError(w, 404, "Not Found")
}
