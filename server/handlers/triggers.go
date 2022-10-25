package handlers

import (
	"Area/database"
	"Area/database/models"
	"Area/lib"
	"bytes"
	"encoding/gob"
	"encoding/json"
	"log"
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
	var buf bytes.Buffer
	var data models.TriggerData
	all := r.URL.Query().Get("a")

	user, err := database.User.GetFromContext(r.Context())
	lib.CheckError(err)
	triggers, _ := database.Trigger.Get(user.ID, all == "true")

	if all == "true" {
		var resp triggersResponse
		copier.Copy(&resp.Triggers, &triggers)
		for i := range triggers {
			buf.Reset()
			buf.Write(triggers[i].Data)
			gob.NewDecoder(&buf).Decode(&data)
			resp.Triggers[i].ActionData = data.ActionData
			resp.Triggers[i].ReactionData = data.ReactionData
		}
		lib.SendJson(w, resp)
	} else {
		var resp triggersSmallResponse
		copier.CopyWithOption(&resp.Triggers, &triggers, copier.Option{DeepCopy: false})
		for i := range triggers {
			buf.Reset()
			buf.Write(triggers[i].Data)
			gob.NewDecoder(&buf).Decode(&data)
			resp.Triggers[i].ActionData = data.ActionData
			resp.Triggers[i].ReactionData = data.ReactionData
		}
		lib.SendJson(w, resp)
	}
}

func CreateTriggers(w http.ResponseWriter, r *http.Request) {
	var input TriggerRequestBody
	var resp triggerSmallResponse
	var data models.Trigger
	var buf bytes.Buffer
	var triggerData models.TriggerData

	user, err := database.User.GetFromContext(r.Context())
	lib.CheckError(err)

	err = json.NewDecoder(r.Body).Decode(&input)
	lib.CheckError(err)

	copier.CopyWithOption(&data, &input, copier.Option{IgnoreEmpty: true})
	data.UserID = user.ID

	trigger, _ := database.Trigger.Create(data)

	buf.Write(trigger.Data)
	gob.NewDecoder(&buf).Decode(&triggerData)
	copier.CopyWithOption(&trigger, &triggerData, copier.Option{IgnoreEmpty: true})

	copier.Copy(&resp.Trigger, &trigger)
	lib.SendJson(w, resp)
}

func GetTriggerById(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	var data models.TriggerData

	all := r.URL.Query().Get("a")
	user, err := database.User.GetFromContext(r.Context())
	lib.CheckError(err)

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	lib.CheckError(err)

	trigger, err := database.Trigger.GetById(uint(id), user.ID, all == "true")
	lib.CheckError(err)

	buf.Write(trigger.Data)
	gob.NewDecoder(&buf).Decode(&data)
	copier.CopyWithOption(&trigger, &data, copier.Option{IgnoreEmpty: true})

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
	var resp triggerSmallResponse
	var buf bytes.Buffer
	var data models.TriggerData

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

	buf.Write(trigger.Data)
	gob.NewDecoder(&buf).Decode(&data)
	copier.CopyWithOption(&data, &input, copier.Option{IgnoreEmpty: true})
	trigger.Data = lib.EncodeToBytes(data)

	trigger, err = database.Trigger.Update(trigger)
	lib.CheckError(err)

	log.Println(trigger.Title)

	copier.Copy(&resp.Trigger, &trigger)
	copier.CopyWithOption(&resp.Trigger, &data, copier.Option{IgnoreEmpty: true})
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
