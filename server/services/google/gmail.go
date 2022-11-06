package google

import (
	"Area/database"
	"Area/database/models"
	"Area/lib"
	"bytes"
	"context"
	"encoding/gob"
	"log"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

func createGoogleConnection(refresh_token string) *gmail.Service {
	var conf = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes: []string{
			"https://www.googleapis.com/auth/google.readonly",
		},
		Endpoint: google.Endpoint,
	}
	var token oauth2.Token = oauth2.Token{
		RefreshToken: refresh_token,
	}

	if refresh_token == "" {
		return nil
	}

	client := conf.Client(context.Background(), &token)
	srv, err := gmail.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return srv
}

func fetchLastGmailReceive(srv *gmail.Service) *gmail.Message {
	res, err := srv.Users.Messages.List("me").Q("label:Inbox").Do()
	if err != nil {
		lib.LogError(err)
		return nil
	}

	id := res.Messages[0].Id
	res2, err := srv.Users.Messages.Get("me", id).Do()
	lib.LogError(err)
	return res2
}

func fetchLastGmailSend(srv *gmail.Service) (*gmail.Message, error) {
	res, err := srv.Users.Messages.List("me").Q("label:Sent").Do()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	id := res.Messages[0].Id
	res2, err := srv.Users.Messages.Get("me", id).Do()
	lib.LogError(err)
	return res2, err
}

func compareGmailData(newData models.TriggerData, oldData models.TriggerData, mail *gmail.Message, trigger *models.Trigger) bool {
	newData.Timestamp = time.UnixMilli(mail.InternalDate)
	if trigger.Data == nil || oldData.Timestamp.Before(newData.Timestamp) {
		newData.Description = mail.Snippet

		for i := range mail.Payload.Headers {
			if mail.Payload.Headers[i].Name == "From" {
				newData.Author = mail.Payload.Headers[i].Value
			}
			if mail.Payload.Headers[i].Name == "Subject" {
				newData.Title = mail.Payload.Headers[i].Value
			}
		}
		newData.ActionData = oldData.ActionData
		newData.ReactionData = oldData.ReactionData

		if trigger.Data != nil {
			trigger.Data = lib.EncodeToBytes(newData)
			database.Trigger.Update(trigger)
			return true
		}
		trigger.Data = lib.EncodeToBytes(newData)
		database.Trigger.Update(trigger)
	}
	return false
}

func checkGmailReceive(srv *gmail.Service, triggerId uint, userId uint) bool {
	var newData models.TriggerData
	var storedData models.TriggerData
	var mail = fetchLastGmailReceive(srv)
	var buf bytes.Buffer

	trigger, err := database.Trigger.GetById(triggerId, userId)
	lib.LogError(err)
	buf.Write(trigger.Data)

	gob.NewDecoder(&buf).Decode(&storedData)

	if mail == nil {
		return false
	}
	return compareGmailData(newData, storedData, mail, trigger)
}

func checkGmailSend(srv *gmail.Service, triggerId uint, userId uint) bool {
	var newData models.TriggerData
	var storedData models.TriggerData
	var mail, err = fetchLastGmailSend(srv)
	var buf bytes.Buffer

	trigger, err := database.Trigger.GetById(triggerId, userId)
	lib.LogError(err)
	buf.Write(trigger.Data)

	gob.NewDecoder(&buf).Decode(&storedData)

	if mail == nil {
		return false
	}
	return compareGmailData(newData, storedData, mail, trigger)
}
