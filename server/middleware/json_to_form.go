package middleware

import (
	"Area/lib"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func JsonToForm(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			body := make(map[string]string)
			form := url.Values{}
			var b []byte
			r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			r.Body.Read(b)
			log.Println(b)
			err := json.NewDecoder(r.Body).Decode(&body)
			lib.CheckError(err)
			for k, v := range body {
				form.Add(k, v)
			}
			r.Body = io.NopCloser(strings.NewReader(form.Encode()))
			r.ContentLength = int64(len(form))
			next.ServeHTTP(w, r)
		}
	})
}
