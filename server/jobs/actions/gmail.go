package actions

import (
	"Area/database/models"
	"context"
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
	res, err := srv.Users.Messages.List("me").Do()
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

func checkReceive(srv *gmail.Service, triggerId uint) bool {
	var data models.TriggerData
	var mail, err = fetchLastMail(srv)

	if err == nil {
		data.Timestamp = time.UnixMilli(mail.InternalDate)
		data.Description = mail.Snippet
		for i := range mail.Payload.Headers {
			if mail.Payload.Headers[i].Name == "From" {
				data.Author = mail.Payload.Headers[i].Value
			}
			if mail.Payload.Headers[i].Name == "Subject" {
				data.Title = mail.Payload.Headers[i].Value
			}
		}
		// check if stored timestamp is smaller than lastMail timestamp
		// if yes then store new mail in db
		// if yes return true
		// else return false
	}
	return false
}

func CheckGmailAction(action models.Action, user models.User) bool {
	var srv = createGmailService(user.GoogleToken)
	if srv == nil {
		return false
	}
	switch action.Event {
	case "receive":
		return checkReceive(srv, action.TriggerID)
	}
	return false
}
