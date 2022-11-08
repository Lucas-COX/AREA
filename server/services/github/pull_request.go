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
	Title     string
	CreatedAt string
	Author    struct {
		Login string
	}
	Repository struct {
		Name string
	}
}

type pullRequestQuery struct {
	Viewer struct {
		PullRequests struct {
			Edges []struct {
				Node pullRequest
			}
		} `graphql:"pullRequests(first: 10, orderBy: {field: CREATED_AT, direction: DESC})"`
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
		newData.Timestamp = prTimestamp
		newData.Title = pr.Title
		newData.Author = "Opened by " + pr.Author.Login
		newData.Description = "Opened in " + pr.Repository.Name
		trigger.Data = lib.EncodeToBytes(newData)
		database.Trigger.Update(trigger)
		return true
	}
	return false
}

func checkNewPullRequest(srv *githubv4.Client, triggerId uint, userId uint) bool {
	var query pullRequestQuery
	var oldData models.TriggerData
	trigger, _ := database.Trigger.GetById(triggerId, userId)
	err := srv.Query(context.Background(), &query, nil)
	if err != nil {
		lib.LogError(err)
		return false
	}
	gob.NewDecoder(bytes.NewReader(trigger.Data)).Decode(&oldData)
	return comparePullRequestData(query.Viewer.PullRequests.Edges[0].Node, trigger)
}
