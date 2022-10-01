package handlers

import (
	"Area/database"
	"Area/database/models"
	"Area/lib"
	"errors"

	"encoding/json"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var input AuthRequestBody
	var resp AuthResponseBody
	var mysqlErr *mysql.MySQLError

	err := json.NewDecoder(r.Body).Decode(&input)
	lib.CheckError(err)

	password, err := lib.HashPassword(input.Password)
	lib.CheckError(err)
	user, err := database.User.Create(models.User{
		Username: input.Username,
		Password: password,
	})
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		lib.SendError(w, http.StatusBadRequest, "User already exists")
		return
	}
	lib.CheckError(err)
	resp.Token, err = lib.CreateToken(map[string]interface{}{
		"username": user.Username,
		"id":       user.ID,
	})
	lib.CheckError(err)
	lib.SendJson(w, resp)
}
