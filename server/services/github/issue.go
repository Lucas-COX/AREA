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

func compareIssueData(issue issue, trigger *models.Trigger) bool {
	var newData models.TriggerData
	var oldData models.TriggerData
	gob.NewDecoder(bytes.NewReader(trigger.Data)).Decode(&oldData)
	issueTimestamp, _ := time.Parse(time.RFC3339, issue.CreatedAt)
	if oldData.Timestamp.Before(issueTimestamp) {
		newData = oldData
		newData.Timestamp = issueTimestamp
		newData.Title = issue.Title
		newData.Author = "Opened by " + issue.Author.Login + " in " + issue.Repository.Name
		newData.Description = issue.Body
		trigger.Data = lib.EncodeToBytes(newData)
		database.Trigger.Update(trigger)
		return true
	}
	return false
}

func compareClosedIssueData(issue closedIssue, trigger *models.Trigger) bool {
	var newData models.TriggerData
	var oldData models.TriggerData
	gob.NewDecoder(bytes.NewReader(trigger.Data)).Decode(&oldData)
	issueTimestamp, _ := time.Parse(time.RFC3339, issue.ClosedAt)
	if oldData.Timestamp.Before(issueTimestamp) {
		newData = oldData
		newData.Timestamp = issueTimestamp
		newData.Title = issue.Title
		newData.Author = "Was opened by " + issue.Author.Login + " in " + issue.Repository.Name
		newData.Description = issue.Body
		trigger.Data = lib.EncodeToBytes(newData)
		database.Trigger.Update(trigger)
		return true
	}
	return false
}

func checkNewIssue(srv *githubv4.Client, triggerId uint, userId uint) bool {
	var query issuesQuery
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
	if len(query.Repository.Issues.Edges) == 0 {
		return false
	}
	return compareIssueData(query.Repository.Issues.Edges[0].Node, trigger)
}

func createIssue(srv *githubv4.Client, storedData models.TriggerData, action string, service string) {
	var mutation createIssueMutation
	var repository = getRepository(srv, storedData.ReactionData)
	var input = githubv4.CreateIssueInput{
		RepositoryID: repository.Viewer.Repository.Id,
		Title:        githubv4.String(storedData.Title),
		Body:         githubv4.NewString(githubv4.String(storedData.Description + "\n\n" + "Created by " + storedData.Author)),
	}
	err := srv.Mutate(context.Background(), &mutation, input, nil)
	lib.LogError(err)
}

func checkClosedIssue(srv *githubv4.Client, trggerId uint, userId uint) bool {
	var query closedIssuesQuery
	var oldData models.TriggerData
	trigger, _ := database.Trigger.GetById(trggerId, userId)
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
	if len(query.Repository.Issues.Edges) == 0 {
		return false
	}
	return compareClosedIssueData(query.Repository.Issues.Edges[0].Node, trigger)
}
