package handlers

import (
	"Area/database"
	"Area/lib"
	"errors"

	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var input AuthRequestBody
	var resp AuthResponseBody

	err := json.NewDecoder(r.Body).Decode(&input)
	lib.CheckError(err)

	user, err := database.User.GetByUsername(input.Username, false)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		lib.SendError(w, http.StatusBadRequest, "User not found")
		return
	}
	lib.CheckError(err)
	if !lib.CheckPassword(input.Password, user.Password) {
		lib.SendError(w, http.StatusBadRequest, "Invalid credentials")
		return
	}

	resp.Token, err = lib.CreateToken(map[string]interface{}{
		"username": user.Username,
		"id":       user.ID,
	})
	lib.CheckError(err)
	lib.SendJson(w, resp)
}
