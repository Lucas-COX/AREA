package github

import (
	"context"

	"github.com/shurcooL/githubv4"
)

func getUser(srv *githubv4.Client) *userQuery {
	var user userQuery
	err := srv.Query(context.Background(), &user, nil)
	if err != nil {
		return nil
	}
	return &user
}

func getRepository(srv *githubv4.Client, name string) *repositoryQuery {
	var repo repositoryQuery
	err := srv.Query(context.Background(), &repo, map[string]interface{}{
		"name": githubv4.String(name),
	})
	if err != nil {
		return nil
	}
	return &repo
}
