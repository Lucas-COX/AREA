package middleware

import (
	"Area/lib"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func JsonToForm(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			body := make(map[string]string)
			form := ""
			i := 0
			r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			err := json.NewDecoder(r.Body).Decode(&body)
			lib.CheckError(err)
			for k, v := range body {
				form += fmt.Sprintf("%s=%s", k, v)
				if i != len(body)-1 {
					form += "&"
				}
				i++
			}
			r.Body = io.NopCloser(strings.NewReader(form))
			r.ContentLength = int64(len(form))
			next.ServeHTTP(w, r)
		}
	})
}
