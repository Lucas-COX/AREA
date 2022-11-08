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
