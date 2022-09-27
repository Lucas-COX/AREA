package handlers

import (
	"Area/lib"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	lib.SendJson(w, ErrorBody{
		Message: "Not found",
	})
}
