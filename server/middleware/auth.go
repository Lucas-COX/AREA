package middleware

import (
	"net/http"

	"github.com/go-chi/oauth"
)

type UserVerifier struct {
}

func (u *UserVerifier) ValidateUser(username, password, scope string, r *http.Request) error {
	return nil
}

func (u *UserVerifier) ValidateClient(clientId, clientSecret, scope string, r *http.Request) error {
	return nil
}

func (u *UserVerifier) ValidateCode(clientId, clientSecret, code, redirectURI string, r *http.Request) (string, error) {
	return "", nil
}

func (u *UserVerifier) AddClaims(tokenType oauth.TokenType, credential, tokenId, scope string, r *http.Request) (map[string]string, error) {
	claims := make(map[string]string)
	return claims, nil
}

func (u *UserVerifier) AddProperties(tokenType oauth.TokenType, credential, tokenID, scope string, r *http.Request) (map[string]string, error) {
	properties := make(map[string]string)
	return properties, nil
}

func (u *UserVerifier) ValidateTokenID(tokenType oauth.TokenType, credential, tokenID, refreshTokenID string) error {
	return nil
}

func (u *UserVerifier) StoreTokenID(tokenType oauth.TokenType, credential, tokenID, refreshTokenID string) error {
	return nil
}
