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
)

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

func compareMergedPullRequestData(pr mergedPullRequest, trigger *models.Trigger) bool {
	var newData models.TriggerData
	var oldData models.TriggerData
	gob.NewDecoder(bytes.NewReader(trigger.Data)).Decode(&oldData)
	prTimestamp, _ := time.Parse(time.RFC3339, pr.MergedAt)
	if oldData.Timestamp.Before(prTimestamp) {
		newData = oldData
		newData.Timestamp = prTimestamp.Local()
		newData.Title = pr.Title
		newData.Author = "Opened by " + pr.MergedBy.Login + " in " + pr.Repository.Name
		newData.Description = "From " + pr.HeadRefName + " to " + pr.BaseRefName
		trigger.Data = lib.EncodeToBytes(newData)
		database.Trigger.Update(trigger)
		return true
	}
	return false
}

func checkNewPullRequest(srv *githubv4.Client, triggerId uint, userId uint) bool {
	var query pullRequestsQuery
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
	if len(query.Repository.PullRequests.Edges) == 0 {
		return false
	}
	return comparePullRequestData(query.Repository.PullRequests.Edges[0].Node, trigger)
}

func checkMergedPullRequest(srv *githubv4.Client, triggerId uint, userId uint) bool {
	var query mergedPullRequestsQuery
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
	if len(query.Repository.PullRequests.Edges) == 0 {
		return false
	}
	return compareMergedPullRequestData(query.Repository.PullRequests.Edges[0].Node, trigger)
}
