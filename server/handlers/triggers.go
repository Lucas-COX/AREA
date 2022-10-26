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

func GetTriggers(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	var data models.TriggerData

	user, err := database.User.GetFromContext(r.Context())
	lib.CheckError(err)
	triggers, _ := database.Trigger.Get(user.ID)

	var resp triggersResponse
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

func CreateTriggers(w http.ResponseWriter, r *http.Request) {
	var input TriggerRequestBody
	var resp triggerResponse
	var data models.Trigger
	var buf bytes.Buffer
	var triggerData models.TriggerData

	user, err := database.User.GetFromContext(r.Context())
	lib.CheckError(err)

	err = json.NewDecoder(r.Body).Decode(&input)
	lib.CheckError(err)

	copier.CopyWithOption(&data, &input, copier.Option{IgnoreEmpty: true})
	triggerData.ActionData = input.ActionData
	triggerData.ReactionData = input.ReactionData

	data.UserID = user.ID
	data.Data = lib.EncodeToBytes(triggerData)

	trigger, _ := database.Trigger.Create(data)

	buf.Write(trigger.Data)
	gob.NewDecoder(&buf).Decode(&triggerData)
	resp.Trigger.ActionData = triggerData.ActionData
	resp.Trigger.ReactionData = triggerData.ReactionData

	copier.Copy(&resp.Trigger, &trigger)
	lib.SendJson(w, resp)
}

func GetTriggerById(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	var data models.TriggerData

	user, err := database.User.GetFromContext(r.Context())
	lib.CheckError(err)

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	lib.CheckError(err)

	trigger, err := database.Trigger.GetById(uint(id), user.ID)
	lib.CheckError(err)

	buf.Write(trigger.Data)
	gob.NewDecoder(&buf).Decode(&data)

	var resp triggerResponse
	copier.CopyWithOption(&resp.Trigger, &trigger, copier.Option{DeepCopy: false})
	resp.Trigger.ActionData = data.ActionData
	resp.Trigger.ReactionData = data.ReactionData
	lib.SendJson(w, resp)
}

func UpdateTrigger(w http.ResponseWriter, r *http.Request) {
	var input TriggerRequestBody
	var resp triggerResponse
	var buf bytes.Buffer
	var data models.TriggerData

	user, err := database.User.GetFromContext(r.Context())
	lib.CheckError(err)

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	lib.CheckError(err)

	err = json.NewDecoder(r.Body).Decode(&input)
	lib.CheckError(err)

	trigger, err := database.Trigger.GetById(uint(id), user.ID)
	lib.CheckError(err)

	copier.CopyWithOption(&trigger, &input, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if input.Active != nil && !*input.Active {
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
	resp.Trigger.ActionData = data.ActionData
	resp.Trigger.ReactionData = data.ReactionData
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
