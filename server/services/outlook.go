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
	"golang.org/x/oauth2/microsoft"
)

type outlookService struct {
	actions   []Action
	reactions []Reaction
}

func (*outlookService) Authenticate(callback string, userId uint) string {
	var state OauthState

	state.Callback = callback
	state.UserId = userId

	conf := &oauth2.Config{
		ClientID:     os.Getenv("MICROSOFT_CLIENT_ID"),
		ClientSecret: os.Getenv("MICROSOFT_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("OAUTH_REDIRECT_URL") + "/providers/microsoft/callback",
		Scopes: []string{
			"https://graph.microsoft.com/Mail.Read",
			"https://graph.microsoft.com/User.Read",
			"offline_access",
		},
		Endpoint: microsoft.AzureADEndpoint("consumers"),
	}
	bytes, _ := json.Marshal(state)
	str := base64.RawStdEncoding.EncodeToString(bytes)
	return conf.AuthCodeURL(str, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
}

func (*outlookService) AuthenticateCallback(base64State string, code string) (string, error) {
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
		ClientID:     os.Getenv("MICROSOFT_CLIENT_ID"),
		ClientSecret: os.Getenv("MICROSOFT_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("OAUTH_REDIRECT_URL") + "/providers/microsoft/callback",
		Scopes: []string{
			"https://graph.microsoft.com/Mail.Read",
			"https://graph.microsoft.com/User.Read",
			"offline_access",
		},
		Endpoint: microsoft.AzureADEndpoint("consumers"),
	}
	token, err := conf.Exchange(context.Background(), code)
	lib.CheckError(err)
	user.MicrosoftToken = token.RefreshToken
	database.User.Update(*user)
	return state.Callback, nil
}

func (outlook *outlookService) GetActions() []Action {
	return outlook.actions
}

func (outlook *outlookService) GetReactions() []Reaction {
	return outlook.reactions
}

func (outlook *outlookService) GetName() string {
	return "outlook"
}

func (*outlookService) Check(action string, trigger models.Trigger) bool {
	var srv = actions.CreateOutlookConnection(trigger.User.MicrosoftToken)
	if srv == nil {
		return false
	}
	switch action {
	case "receive":
		return actions.OutlookReceive(srv, trigger.ID, trigger.UserID)
	}
	return false
}

func (*outlookService) React(reaction string, trigger models.Trigger) {
}

func (outlook *outlookService) ToJson() JsonService {
	return JsonService{
		Name:      outlook.GetName(),
		Actions:   outlook.GetActions(),
		Reactions: outlook.GetReactions(),
	}
}

func NewOutlookService() *outlookService {
	return &outlookService{
		actions: []Action{
			{Name: "receive", Description: "When the user receives an email"},
		},
		reactions: []Reaction{},
	}
}
