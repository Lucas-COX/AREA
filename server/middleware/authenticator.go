package middleware

import (
	"Area/database"
	"Area/lib"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/jwt"
)

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, claims, err := jwtauth.FromContext(r.Context())

		if err != nil {
			lib.SendError(w, http.StatusUnauthorized, err.Error())
			return
		}

		if token == nil || jwt.Validate(token) != nil {
			lib.SendError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			return
		}

		_, err = database.User.GetById(uint(claims["id"].(float64)), false)
		if err != nil {
			lib.SendError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			return
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}
