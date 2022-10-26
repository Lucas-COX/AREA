package services

import (
	"Area/database"
	"Area/database/models"
	"Area/services/actions"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GmailService struct {
	Actions   []Action
	Reactions []Reaction
}

func (gmail *GmailService) Authenticate(redirect string, callback string, userId uint) string {
	var state OauthState

	state.Callback = callback
	state.UserId = userId

	conf := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  redirect,
		Scopes: []string{
			"https://www.googleapis.com/auth/gmail.readonly",
		},
		Endpoint: google.Endpoint,
	}
	bytes, _ := json.Marshal(state)
	str := base64.RawStdEncoding.EncodeToString(bytes)
	return conf.AuthCodeURL(str, oauth2.AccessTypeOffline)
}

func (gmail *GmailService) AuthenticateCallback(base64State string, code string) (string, error) {
	var state OauthState

	bytes, _ := base64.RawStdEncoding.DecodeString(base64State)
	err := json.Unmarshal(bytes, &state)
	if err != nil {
		return "", errors.New("invalid callback url")
	}
	user, err := database.User.GetById(state.UserId, false)
	if err != nil {
		return "", errors.New("invalid user id")
	}

	conf := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8080/providers/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/gmail.readonly",
		},
		Endpoint: google.Endpoint,
	}
	token, _ := conf.Exchange(context.Background(), code)
	user.GoogleToken = token.RefreshToken
	database.User.Update(*user)
	return state.Callback, nil
}

func (gmail *GmailService) GetActions() []Action {
	return gmail.Actions
}

func (gmail *GmailService) GetReactions() []Reaction {
	return gmail.Reactions
}

func (gmail *GmailService) GetName() string {
	return "gmail"
}

func (gmail *GmailService) Check(action string, trigger models.Trigger) bool {
	var srv = actions.CreateGmailConnection(trigger.User.GoogleToken)
	if srv == nil {
		return false
	}
	switch action {
	case "receive":
		return actions.GmailReceive(srv, trigger.ID, trigger.UserID)
	case "send":
		return actions.GmailSend(srv, trigger.ID, trigger.UserID)
	}
	return false
}

func (gmail *GmailService) React(reaction string, trigger models.Trigger) {
}

func NewGmailService() *GmailService {
	return &GmailService{
		Actions: []Action{
			{Name: "receive", Description: "When the user receives an email"},
			{Name: "send", Description: "When the user sends an email"},
		},
		Reactions: []Reaction{},
	}
}
