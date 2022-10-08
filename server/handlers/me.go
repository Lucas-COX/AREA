package handlers

import (
	"Area/database"
	"Area/lib"
	"net/http"

	"github.com/jinzhu/copier"
)

type userBody struct {
	User UserBody `json:"me"`
}

func Me(w http.ResponseWriter, r *http.Request) {
	var resp userBody
	user, err := database.User.GetFromContext(r.Context())
	lib.CheckError(err)

	copier.Copy(&resp.User, &user)
	lib.SendJson(w, resp)
}
