package handlers

import (
	"Area/database"
	"Area/lib"
	"encoding/json"
	"errors"
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
	lib.SetCookie(w, "area_token", resp.Token)
	lib.SendJson(w, resp)
}

func LoginDone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Area</title>
    <style>
        .container {
            width: 350px;
    		height: 100px;
    		position: absolute;
			top:0;
			bottom: 0;
			left: 0;
			right: 0;

			margin: auto;
        }
    </style>
</head>
<body>
    <p class="container">Successfully connected, you can close this page.</p>
</body>
</html>
`))
}
