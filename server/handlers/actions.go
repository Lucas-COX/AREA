package handlers

import (
	"Area/database"
	"Area/lib"
	"net/http"

	"github.com/jinzhu/copier"
)

type actionsResponse struct {
	Actions []ActionResponseBody `json:"actions"`
}

func GetActions(w http.ResponseWriter, r *http.Request) {
	var resp actionsResponse

	_, err := database.User.GetFromContext(r.Context())
	lib.CheckError(err)

	actions, _ := database.Action.Get()

	copier.Copy(&resp.Actions, &actions)
	lib.SendJson(w, resp)
}
