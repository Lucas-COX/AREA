package github

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

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type githubService struct {
	actions   []types.Action
	reactions []types.Reaction
}

func createGithubConnection(refresh_token string) *githubv4.Client {
	token := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: refresh_token},
	)
	if refresh_token == "" {
		return nil
	}
	client := oauth2.NewClient(context.Background(), token)
	return githubv4.NewClient(client)
}

func (*githubService) Authenticate(callback string, userId uint) string {
	var state types.OauthState

	state.Callback = callback
	state.UserId = userId

	conf := &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("OAUTH_REDIRECT_URL") + "/providers/github/callback",
		Scopes: []string{
			"access_offline",
			"repo",
		},
		Endpoint: github.Endpoint,
	}
	bytes, _ := json.Marshal(state)
	str := base64.RawStdEncoding.EncodeToString(bytes)
	return conf.AuthCodeURL(str, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
}

func (*githubService) AuthenticateCallback(base64State string, code string) (string, error) {
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
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("OAUTH_REDIRECT_URL") + "/providers/github/callback",
		Scopes: []string{
			"access_offline",
		},
		Endpoint: github.Endpoint,
	}
	token, err := conf.Exchange(context.Background(), code)
	lib.CheckError(err)
	user.GithubToken = token.AccessToken
	database.User.Update(*user)
	return state.Callback, nil
}

func (gh *githubService) GetActions() []types.Action {
	return gh.actions
}

func (gh *githubService) GetReactions() []types.Reaction {
	return gh.reactions
}

func (gh *githubService) GetName() string {
	return "github"
}

func (*githubService) Check(action string, trigger models.Trigger) bool {
	var srv = createGithubConnection(trigger.User.GithubToken)
	if srv == nil {
		return false
	}
	switch action {
	case "pull request":
		return checkNewPullRequest(srv, trigger.ID, trigger.UserID)
	case "issue":
		return checkNewIssue(srv, trigger.ID, trigger.UserID)
	}
	return false
}

func (*githubService) React(reaction string, trigger models.Trigger) {
}

func (gh *githubService) ToJson() types.JsonService {
	return types.JsonService{
		Name:      gh.GetName(),
		Actions:   gh.GetActions(),
		Reactions: gh.GetReactions(),
	}
}

func New() *githubService {
	return &githubService{
		actions: []types.Action{
			{Name: "pull request", Description: "When a pull request is opened"},
		},
		reactions: []types.Reaction{},
	}
}
