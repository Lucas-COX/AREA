package github

import (
	"Area/database"
	"Area/database/models"
	"Area/lib"
	"bytes"
	"context"
	"encoding/gob"
	"time"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type pullRequest struct {
	Title       string
	BaseRefName string
	HeadRefName string
	CreatedAt   string
	Author      struct {
		Login string
	}
	Repository struct {
		Name string
	}
}

type repositoryQuery struct {
	Repository struct {
		PullRequests struct {
			Edges []struct {
				Node pullRequest
			}
		} `graphql:"pullRequests(first: 1, orderBy: {field: CREATED_AT, direction: DESC})"`
	} `graphql:"repository(owner: $owner, name: $name)"`
}

type userQuery struct {
	Viewer struct {
		Login string
	}
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

func comparePullRequestData(pr pullRequest, trigger *models.Trigger) bool {
	var newData models.TriggerData
	var oldData models.TriggerData
	gob.NewDecoder(bytes.NewReader(trigger.Data)).Decode(&oldData)
	prTimestamp, _ := time.Parse(time.RFC3339, pr.CreatedAt)
	if oldData.Timestamp.Before(prTimestamp) {
		newData = oldData
		newData.Timestamp = prTimestamp.Local()
		newData.Title = pr.Title
		newData.Author = "Opened by " + pr.Author.Login + " in " + pr.Repository.Name
		newData.Description = "From " + pr.HeadRefName + " to " + pr.BaseRefName
		trigger.Data = lib.EncodeToBytes(newData)
		database.Trigger.Update(trigger)
		return true
	}
	return false
}

func getUser(srv *githubv4.Client) *userQuery {
	var user userQuery
	err := srv.Query(context.Background(), &user, nil)
	if err != nil {
		return nil
	}
	return &user
}

func checkNewPullRequest(srv *githubv4.Client, triggerId uint, userId uint) bool {
	var query repositoryQuery
	var oldData models.TriggerData
	trigger, _ := database.Trigger.GetById(triggerId, userId)
	gob.NewDecoder(bytes.NewReader(trigger.Data)).Decode(&oldData)

	user := getUser(srv)
	err := srv.Query(context.Background(), &query, map[string]interface{}{
		"owner": githubv4.String(user.Viewer.Login),
		"name":  githubv4.String(oldData.ActionData),
	})
	if err != nil {
		lib.LogError(err)
		return false
	}
	gob.NewDecoder(bytes.NewReader(trigger.Data)).Decode(&oldData)
	return comparePullRequestData(query.Repository.PullRequests.Edges[0].Node, trigger)
}
