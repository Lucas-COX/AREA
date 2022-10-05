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

func GetTriggers(w http.ResponseWriter, r *http.Request) {
	var resp triggersResponse
	user, err := UserFromContext(r.Context())
	lib.CheckError(err)

	copier.Copy(&resp.Triggers, &user.Triggers)
	lib.SendJson(w, resp)
}

func CreateTriggers(w http.ResponseWriter, r *http.Request) {
	var input TriggerRequestBody
	var resp triggerResponse
	user, err := UserFromContext(r.Context())
	lib.CheckError(err)
	err = json.NewDecoder(r.Body).Decode(&input)
	lib.CheckError(err)
	trigger, _ := database.Trigger.Create(models.Trigger{
		Title:       input.Title,
		Description: input.Description,
		UserID:      user.ID,
	})
	copier.Copy(&resp.Trigger, &trigger)
	lib.SendJson(w, resp)
}

func GetTriggerById(w http.ResponseWriter, r *http.Request) {
	var resp triggerResponse
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	lib.CheckError(err)

	trigger, err := database.Trigger.GetById(uint(id))
	lib.CheckError(err)

	copier.Copy(&resp.Trigger, &trigger)
	lib.SendJson(w, resp)
}

func UpdateTrigger(w http.ResponseWriter, r *http.Request) {
	var input TriggerRequestBody
	var resp triggerResponse
	user, err := UserFromContext(r.Context())
	lib.CheckError(err)
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	lib.CheckError(err)
	err = json.NewDecoder(r.Body).Decode(&input)
	lib.CheckError(err)

	trigger, err := database.Trigger.GetById(uint(id))
	lib.CheckError(err)
	if user.ID != trigger.UserID {
		lib.SendError(w, http.StatusUnauthorized, "Can't Modify this trigger")
		return
	}

	if input.Title != "" {
		trigger.Title = input.Title
	}
	if input.Description != "" {
		trigger.Description = input.Description
	}
	trigger, err = database.Trigger.Update(trigger)
	lib.CheckError(err)

	copier.Copy(&resp.Trigger, &trigger)
	lib.SendJson(w, resp)
}

func DeleteTrigger(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	lib.CheckError(err)

	user, err := UserFromContext(r.Context())
	lib.CheckError(err)

	trigger, err := database.Trigger.GetById(uint(id))
	lib.CheckError(err)

	if user.ID != trigger.UserID {
		lib.SendError(w, http.StatusUnauthorized, "Can't delete this trigger")
		return
	}

	err = database.Trigger.Delete(trigger)
	lib.CheckError(err)

	lib.SendJson(w, "Trigger deleted")
}
