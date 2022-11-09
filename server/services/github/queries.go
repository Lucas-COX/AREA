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

type mergedPullRequest struct {
	Title    string
	MergedAt string
	MergedBy struct {
		Login string
	}
	BaseRefName string
	HeadRefName string
	Repository  struct {
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

type closedIssue struct {
	Title    string
	Body     string
	ClosedAt string
	Author   struct {
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

type mergedPullRequestsQuery struct {
	Repository struct {
		PullRequests struct {
			Edges []struct {
				Node mergedPullRequest
			}
		} `graphql:"pullRequests(states: MERGED, first: 1, orderBy: {field: CREATED_AT, direction: DESC})"`
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

type closedIssuesQuery struct {
	Repository struct {
		Issues struct {
			Edges []struct {
				Node closedIssue
			}
		} `graphql:"issues(states: CLOSED, first: 1, orderBy: {field: UPDATED_AT, direction: DESC})"`
	} `graphql:"repository(owner: $owner, name: $name)"`
}

type userQuery struct {
	Viewer struct {
		Login string
	}
}

type repositoryQuery struct {
	Viewer struct {
		Login      string
		Repository struct {
			Name string
			Id   string
		} `graphql:"repository(name: $name)"`
	}
}

type commit struct {
	CommittedDate string
	Message       string
	MessageBody   string
	Author        struct {
		Name  string
		Email string
	}
	Repository struct {
		Name string
	}
	CommitUrl string
}

type commitQuery struct {
	Repository struct {
		Refs struct {
			Edges []struct {
				Node struct {
					Name   string
					Target struct {
						Commit struct {
							History struct {
								Edges []struct {
									Node struct {
										Commit commit `graphql:"... on Commit"`
									}
								}
							} `graphql:"history(first: 1)"`
						} `graphql:"... on Commit"`
					}
				}
			}
		} `graphql:"refs(refPrefix: \"refs/heads/\", orderBy: {direction: DESC, field: TAG_COMMIT_DATE}, first: 5)"`
	} `graphql:"repository(owner: $owner, name: $name)"`
}

type createIssueMutation struct {
	CreateIssue struct {
		ClientMutationId string
	} `graphql:"createIssue(input: $input)"`
}
