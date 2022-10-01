package lib

import (
	"os"
	"strconv"

	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"
)

var tokenAuth *jwtauth.JWTAuth

func NewTokenAuth() *jwtauth.JWTAuth {
	tokenAuth = jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET")), nil)
	return tokenAuth
}

func CreateToken(claims map[string]interface{}) (string, error) {
	_, token, err := tokenAuth.Encode(claims)
	return token, err
}

func HashPassword(password string) (string, error) {
	salt, err := strconv.Atoi(os.Getenv("SALT"))
	if err != nil {
		return "", err
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(bytes), err
}

func CheckPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
