package handlers

import (
	"Area/database"
	"Area/database/models"
	"Area/lib"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi/v5"
	"github.com/jinzhu/copier"
)

type triggersResponse struct {
	Triggers []TriggerBody `json:"triggers"`
}

type triggerResponse struct {
	Trigger TriggerBody `json:"trigger"`
}

func GetTriggers(w http.ResponseWriter, r *http.Request) {
	var resp triggersResponse
<<<<<<< HEAD
<<<<<<< HEAD

	user, err := database.User.GetFromContext(r.Context())
	lib.CheckError(err)
	triggers, _ := database.Trigger.Get(user.ID)

=======
	user, err := UserFromContext(r.Context())
=======

	user, err := database.User.GetFromContext(r.Context())
>>>>>>> 7b575bd (feat(server): implement google oauth flow to get refresh token)
	lib.CheckError(err)
	triggers, _ := database.Trigger.Get(user.ID)

>>>>>>> ae11fa4 (feat(server): add actions and reactions to trigger model)
	copier.Copy(&resp.Triggers, &triggers)
	lib.SendJson(w, resp)
}

func CreateTriggers(w http.ResponseWriter, r *http.Request) {
	var input TriggerRequestBody
	var resp triggerResponse
	var data models.Trigger

	user, err := database.User.GetFromContext(r.Context())
	lib.CheckError(err)

	err = json.NewDecoder(r.Body).Decode(&input)
	lib.CheckError(err)

	copier.Copy(&data, &input)
	data.UserID = user.ID

	trigger, _ := database.Trigger.Create(data)

	copier.Copy(&resp.Trigger, &trigger)
	lib.SendJson(w, resp)
}

func GetTriggerById(w http.ResponseWriter, r *http.Request) {
	var resp triggerResponse

<<<<<<< HEAD
<<<<<<< HEAD
	user, err := database.User.GetFromContext(r.Context())
=======
	user, err := UserFromContext(r.Context())
>>>>>>> ae11fa4 (feat(server): add actions and reactions to trigger model)
=======
	user, err := database.User.GetFromContext(r.Context())
>>>>>>> 7b575bd (feat(server): implement google oauth flow to get refresh token)
	lib.CheckError(err)

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	lib.CheckError(err)

	trigger, err := database.Trigger.GetById(uint(id), user.ID)
	lib.CheckError(err)

	copier.Copy(&resp.Trigger, &trigger)
	lib.SendJson(w, resp)
}

func UpdateTrigger(w http.ResponseWriter, r *http.Request) {
	var input TriggerRequestBody
	var resp triggerResponse
	user, err := database.User.GetFromContext(r.Context())
	lib.CheckError(err)

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	lib.CheckError(err)

	err = json.NewDecoder(r.Body).Decode(&input)
	lib.CheckError(err)

	trigger, err := database.Trigger.GetById(uint(id), user.ID)
	lib.CheckError(err)

	copier.CopyWithOption(&trigger, &input, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	spew.Dump(trigger)
	trigger, err = database.Trigger.Update(trigger)
	lib.CheckError(err)

	copier.Copy(&resp.Trigger, trigger)
	lib.SendJson(w, resp)
}

func DeleteTrigger(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	lib.CheckError(err)

	user, err := database.User.GetFromContext(r.Context())
	lib.CheckError(err)

	trigger, err := database.Trigger.GetById(uint(id), user.ID)
	lib.CheckError(err)
	if user.ID != trigger.UserID {
		lib.SendError(w, http.StatusUnauthorized, "Can't delete this trigger")
		return
	}

	err = database.Trigger.Delete(trigger)
	lib.CheckError(err)

	lib.SendJson(w, "Trigger deleted")
}
