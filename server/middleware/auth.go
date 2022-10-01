package middleware

import (
	"Area/database"
	"net/http"

	"github.com/go-chi/oauth"
)

type UserVerifier struct {
}

// For user / password authentication
func (u *UserVerifier) ValidateUser(username, password, scope string, r *http.Request) error {
	_, err := database.User.GetByUsername(username)
	return err
}

// For client authentication
func (u *UserVerifier) ValidateClient(clientId, clientSecret, scope string, r *http.Request) error {
	// Todo: check in db and validate if clientId and clientSecret match ?
	return nil
}

// UNUSED
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

// For refresh token
func (u *UserVerifier) ValidateTokenID(tokenType oauth.TokenType, credential, tokenID, refreshTokenID string) error {
	// Todo: check in the users table for credentials and get the stored token ID
	return nil
}

// For refresh token
func (u *UserVerifier) StoreTokenID(tokenType oauth.TokenType, credential, tokenID, refreshTokenID string) error {
	// Todo: add the token to a column in the users table to keep it and use it in ValidateTokenID
	return nil
}
