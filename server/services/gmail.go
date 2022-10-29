package services

import (
	"Area/database"
	"Area/database/models"
	"Area/lib"
	"Area/services/actions"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type gmailService struct {
	actions   []Action
	reactions []Reaction
}

func (*gmailService) Authenticate(callback string, userId uint) string {
	var state OauthState

	state.Callback = callback
	state.UserId = userId

	conf := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("OAUTH_REDIRECT_URL") + "/providers/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/gmail.readonly",
		},
		Endpoint: google.Endpoint,
	}
	bytes, _ := json.Marshal(state)
	str := base64.RawStdEncoding.EncodeToString(bytes)
	return conf.AuthCodeURL(str, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
}

func (*gmailService) AuthenticateCallback(base64State string, code string) (string, error) {
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
		RedirectURL:  os.Getenv("OAUTH_REDIRECT_URL") + "/providers/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/gmail.readonly",
		},
		Endpoint: google.Endpoint,
	}
	token, err := conf.Exchange(context.Background(), code)
	lib.CheckError(err)
	user.GoogleToken = token.RefreshToken
	database.User.Update(*user)
	return state.Callback, nil
}

func (gmail *gmailService) GetActions() []Action {
	return gmail.actions
}

func (gmail *gmailService) GetReactions() []Reaction {
	return gmail.reactions
}

func (gmail *gmailService) GetName() string {
	return "gmail"
}

func (*gmailService) Check(action string, trigger models.Trigger) bool {
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

func (*gmailService) React(reaction string, trigger models.Trigger) {
}

func (gmail *gmailService) ToJson() JsonService {
	return JsonService{
		Name:      gmail.GetName(),
		Actions:   gmail.GetActions(),
		Reactions: gmail.GetReactions(),
	}
}

func NewGmailService() *gmailService {
	return &gmailService{
		actions: []Action{
			{Name: "receive", Description: "When the user receives an email"},
			{Name: "send", Description: "When the user sends an email"},
		},
		reactions: []Reaction{},
	}
}
