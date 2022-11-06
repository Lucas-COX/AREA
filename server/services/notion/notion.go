package notion

import (
	"Area/database"
	"Area/database/models"
	"Area/lib"
	"Area/services/types"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"os"

	"golang.org/x/oauth2"
)

type notionService struct {
	actions   []types.Action
	reactions []types.Reaction
}

func (*notionService) Authenticate(callback string, userId uint) string {
	var state types.OauthState

	state.Callback = callback
	state.UserId = userId

	conf := &oauth2.Config{
		ClientID:     os.Getenv("NOTION_CLIENT_ID"),
		ClientSecret: os.Getenv("NOTION_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("OAUTH_REDIRECT_URL") + "/providers/notion/callback",
		Scopes: []string{
			"https://www.notionapis.com/auth/notion.readonly",
		},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://api.notion.com/v1/oauth/authorize",
			TokenURL: "https://api.notion.com/v1/oauth/token",
		},
	}
	bytes, _ := json.Marshal(state)
	str := base64.RawStdEncoding.EncodeToString(bytes)
	return conf.AuthCodeURL(str, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
}

func (*notionService) AuthenticateCallback(base64State string, code string) (string, error) {
	var state types.OauthState

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
		ClientID:     os.Getenv("NOTION_CLIENT_ID"),
		ClientSecret: os.Getenv("NOTION_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("OAUTH_REDIRECT_URL") + "/providers/notion/callback",
		Scopes:       []string{},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://api.notion.com/v1/oauth/authorize",
			TokenURL: "https://api.notion.com/v1/oauth/token",
		},
	}
	token, err := conf.Exchange(context.Background(), code)
	lib.CheckError(err)
	user.NotionToken = token.AccessToken
	database.User.Update(*user)
	return state.Callback, nil
}

func (notion *notionService) GetActions() []types.Action {
	return notion.actions
}

func (notion *notionService) GetReactions() []types.Reaction {
	return notion.reactions
}

func (notion *notionService) GetName() string {
	return "notion"
}

func (*notionService) Check(action string, trigger models.Trigger) bool {
	return false
}

func (*notionService) React(reaction string, trigger models.Trigger) {
}

func (notion *notionService) ToJson() types.JsonService {
	return types.JsonService{
		Name:      notion.GetName(),
		Actions:   notion.GetActions(),
		Reactions: notion.GetReactions(),
	}
}

func New() *notionService {
	return &notionService{
		actions:   []types.Action{},
		reactions: []types.Reaction{},
	}
}
