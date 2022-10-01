package middleware

import (
	"Area/lib"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/jwt"
)

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, err := jwtauth.FromContext(r.Context())

		if err != nil {
			lib.SendError(w, http.StatusUnauthorized, err.Error())
			return
		}

		if token == nil || jwt.Validate(token) != nil {
			lib.SendError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			return
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}
