package actions

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

func createGmailService(refresh_token string) *gmail.Service {
	var conf = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes: []string{
			"https://www.googleapis.com/auth/gmail.readonly",
		},
		Endpoint: google.Endpoint,
	}
	var token oauth2.Token = oauth2.Token{
		RefreshToken: refresh_token,
	}

	client := conf.Client(context.Background(), &token)
	srv, err := gmail.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return srv
}

func fetchLastMail(srv *gmail.Service) (*gmail.Message, error) {
	res, err := srv.Users.Messages.List("me").Q("label:Inbox").Do()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	id := res.Messages[0].Id
	res2, err := srv.Users.Messages.Get("me", id).Do()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return res2, nil
}

func checkReceive(srv *gmail.Service, triggerId uint, userId uint) bool {
	var newData models.TriggerData
	var storedData models.TriggerData
	var mail, err = fetchLastMail(srv)
	var buf bytes.Buffer

	trigger, err := database.Trigger.GetById(triggerId, userId)
	lib.CheckError(err)
	buf.Write(trigger.Data)

	gob.NewDecoder(&buf).Decode(&storedData)

	newData.Timestamp = time.UnixMilli(mail.InternalDate)
	if trigger.Data == nil || storedData.Timestamp.Before(newData.Timestamp) {
		newData.Description = mail.Snippet
		for i := range mail.Payload.Headers {
			if mail.Payload.Headers[i].Name == "From" {
				newData.Author = mail.Payload.Headers[i].Value
			}
			if mail.Payload.Headers[i].Name == "Subject" {
				newData.Title = mail.Payload.Headers[i].Value
			}
		}

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

func CheckGmailAction(action models.Action, user models.User) bool {
	log.Println("checking gmail....")
	var srv = createGmailService(user.GoogleToken)
	if srv == nil {
		return false
	}
	switch action.Event {
	case "receive":
		return checkReceive(srv, action.TriggerID, user.ID)
	}
	return false
}
