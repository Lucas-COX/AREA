package handlers

import (
	"Area/database"
	"Area/database/models"
	"Area/lib"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jinzhu/copier"
)

type triggersResponse struct {
	Triggers []TriggerBody `json:"triggers"`
}

type triggerResponse struct {
	Trigger TriggerBody `json:"trigger"`
}

type triggersSmallResponse struct {
	Triggers []TriggerSmallBody `json:"triggers"`
}

type triggerSmallResponse struct {
	Trigger TriggerSmallBody `json:"trigger"`
}

func GetTriggers(w http.ResponseWriter, r *http.Request) {
	all := r.URL.Query().Get("a")

	user, err := database.User.GetFromContext(r.Context())
	lib.CheckError(err)
	triggers, _ := database.Trigger.Get(user.ID, all == "true")

	if all == "true" {
		var resp triggersResponse
		copier.Copy(&resp.Triggers, &triggers)
		lib.SendJson(w, resp)
	} else {
		var resp triggersSmallResponse
		copier.CopyWithOption(&resp.Triggers, &triggers, copier.Option{DeepCopy: false})
		lib.SendJson(w, resp)
	}
}

func CreateTriggers(w http.ResponseWriter, r *http.Request) {
	var input TriggerRequestBody
	var resp triggerSmallResponse
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
	all := r.URL.Query().Get("a")
	user, err := database.User.GetFromContext(r.Context())
	lib.CheckError(err)

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	lib.CheckError(err)

	trigger, err := database.Trigger.GetById(uint(id), user.ID, all == "true")
	lib.CheckError(err)

	if all == "true" {
		var resp triggerResponse
		copier.Copy(&resp.Trigger, &trigger)
		lib.SendJson(w, resp)
	} else {
		var resp triggerSmallResponse
		copier.CopyWithOption(&resp.Trigger, &trigger, copier.Option{DeepCopy: false})
		lib.SendJson(w, resp)
	}
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

	trigger, err := database.Trigger.GetById(uint(id), user.ID, false)
	lib.CheckError(err)

	copier.CopyWithOption(&trigger, &input, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if !input.Active {
		trigger.Active = false
	}

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

	trigger, err := database.Trigger.GetById(uint(id), user.ID, false)
	lib.CheckError(err)
	if user.ID != trigger.UserID {
		lib.SendError(w, http.StatusUnauthorized, "Can't delete this trigger")
		return
	}

	err = database.Trigger.Delete(trigger)
	lib.CheckError(err)

	lib.SendJson(w, "Trigger deleted")
}
