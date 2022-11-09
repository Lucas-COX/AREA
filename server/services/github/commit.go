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

func compareCommitData(commit commit, trigger *models.Trigger) bool {
	var newData models.TriggerData
	var oldData models.TriggerData
	gob.NewDecoder(bytes.NewReader(trigger.Data)).Decode(&oldData)
	commitTimestamp, _ := time.Parse(time.RFC3339, commit.CommittedDate)
	if oldData.Timestamp.Before(commitTimestamp) {
		newData = oldData
		newData.Timestamp = commitTimestamp
		newData.Title = commit.Message
		newData.Author = "Authored by " + commit.Author.Name + " (" + commit.Author.Email + ")"
		newData.Description = "Pushed in " + commit.Repository.Name + " on branch " + commit.MessageBody + "\n" + commit.CommitUrl
		trigger.Data = lib.EncodeToBytes(newData)
		database.Trigger.Update(trigger)
		return true
	}
	return false
}

func getLastCommit(query commitQuery) commit {
	var lastCommit commit
	for _, v := range query.Repository.Refs.Edges {
		if len(v.Node.Target.Commit.History.Edges) == 0 {
			continue
		}
		tmp := v.Node.Target.Commit.History.Edges[0].Node.Commit
		tmp.MessageBody = v.Node.Name
		tmpDate, _ := time.Parse(time.RFC3339, tmp.CommittedDate)
		lastDate, _ := time.Parse(time.RFC3339, lastCommit.CommittedDate)
		if tmpDate.After(lastDate) {
			lastCommit = tmp
		}
	}
	return lastCommit
}

func checkNewCommit(srv *githubv4.Client, triggerId uint, userId uint) bool {
	var query commitQuery
	var user userQuery
	var trigger, err = database.Trigger.GetById(triggerId, userId)
	var data models.TriggerData

	err = gob.NewDecoder(bytes.NewReader(trigger.Data)).Decode(&data)
	if err != nil {
		return false
	}
	err = srv.Query(context.Background(), &user, nil)
	if err != nil {
		lib.LogError(err)
		return false
	}
	err = srv.Query(context.Background(), &query, map[string]interface{}{
		"owner": githubv4.String(user.Viewer.Login),
		"name":  githubv4.String(data.ActionData),
	})
	if err != nil {
		lib.LogError(err)
		return false
	}
	commit := getLastCommit(query)
	return compareCommitData(commit, trigger)
}
