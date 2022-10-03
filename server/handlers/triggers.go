package handlers

import (
	"Area/lib"
	"net/http"

	"github.com/jinzhu/copier"
)

type triggers struct {
	Triggers []TriggerBody `json:"triggers"`
}

func Triggers(w http.ResponseWriter, r *http.Request) {
	var resp triggers
	user, err := UserFromContext(r.Context())
	lib.CheckError(err)

	copier.Copy(&resp.Triggers, &user.Triggers)
	lib.SendJson(w, resp)
}
