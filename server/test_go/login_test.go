package handlers

import (
	"Area/router"
	"net/http"
	"testing"
)

func Testlogin(t *testing.T) {
	pop := router.New()
	handlers := login(pop)

	w := &mock_http.ResponseWriter{}
	r := &http.Request{}

	handler(w, r)
	result := w.Body

	if len(result) != 1 {
		t.Errorf("not op")
	}
}
