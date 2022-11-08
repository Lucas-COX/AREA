package github

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

type issue struct {
	Title     string
	Body      string
	CreatedAt string
	Author    struct {
		Login string
	}
	Repository struct {
		Name string
	}
}

type pullRequestsQuery struct {
	Repository struct {
		PullRequests struct {
			Edges []struct {
				Node pullRequest
			}
		} `graphql:"pullRequests(first: 1, orderBy: {field: CREATED_AT, direction: DESC})"`
	} `graphql:"repository(owner: $owner, name: $name)"`
}

type issuesQuery struct {
	Repository struct {
		Issues struct {
			Edges []struct {
				Node issue
			}
		} `graphql:"issues(first: 1, orderBy: {field: CREATED_AT, direction: DESC})"`
	} `graphql:"repository(owner: $owner, name: $name)"`
}

type userQuery struct {
	Viewer struct {
		Login string
	}
}
