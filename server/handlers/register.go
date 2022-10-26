package handlers

import (
	"Area/database"
	"Area/database/models"
	"Area/lib"
	"strings"

	"encoding/json"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var input AuthRequestBody
	var resp AuthResponseBody

	err := json.NewDecoder(r.Body).Decode(&input)
	lib.CheckError(err)

	password, err := lib.HashPassword(input.Password)
	lib.CheckError(err)
	user, err := database.User.Create(models.User{
		Username: input.Username,
		Password: password,
	})
	if err != nil && strings.Contains(err.Error(), "23505") {
		lib.SendError(w, http.StatusBadRequest, "User already exists")
		return
	}

	lib.CheckError(err)
	resp.Token, err = lib.CreateToken(map[string]interface{}{
		"username": user.Username,
		"id":       user.ID,
	})
	lib.CheckError(err)
	lib.SetCookie(w, "area_token", resp.Token)
	lib.SendJson(w, resp)
}
