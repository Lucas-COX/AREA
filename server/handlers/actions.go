package handlers

import (
	"Area/database"
	"Area/database/models"
	"Area/lib"
	"encoding/json"
	"net/http"

	"github.com/jinzhu/copier"
)

type actionsResponse struct {
	Actions []ActionResponseBody `json:"actions"`
}

type actionResponse struct {
	Action ActionResponseBody `json:"action"`
}

func GetActions(w http.ResponseWriter, r *http.Request) {
	var resp actionsResponse

	_, err := database.User.GetFromContext(r.Context())
	lib.CheckError(err)

	actions, _ := database.Action.Get()

	copier.Copy(&resp.Actions, &actions)
	lib.SendJson(w, resp)
}

func CreateAction(w http.ResponseWriter, r *http.Request) {
	var resp actionResponse
	var input ActionRequestBody

	user, err := database.User.GetFromContext(r.Context())
	lib.CheckError(err)

	if user.Role != models.AdminRole {
		lib.SendError(w, 401, "Unauthorized")
		return
	}

	err = json.NewDecoder(r.Body).Decode(&input)
	lib.CheckError(err)

	action, err := database.Action.Create(&models.Action{
		Type:  input.Type,
		Event: input.Event,
	})
	lib.CheckError(err)

	copier.Copy(&resp.Action, action)
	lib.SendJson(w, resp)
}
