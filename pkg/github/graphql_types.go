package github

import (
	"time"

	"github.com/shurcooL/githubv4"
)

type (
	// Author contains information about a GitHub user
	Author struct {
		Login     string
		AvatarURL string
	}

	// Comment contains information about a comment made on the pull request
	Comment struct {
		Author    Author
		URL       string
		UpdatedAt githubv4.DateTime
		BodyText  string
	}

	// PullRequestCommit contains information about a commit which belongs to the pull request
	PullRequestCommit struct {
		Commit Commit
	}

	StatusContext struct {
		Context string
		State string
	}

	// Commit contains information about a commit from the version control
	Commit struct {
		PushedDate time.Time
		Status     struct {
			State string
			Contexts []StatusContext
		}
	}

	// PullRequest contains information about a pull request
	PullRequest struct {
		Title        string
		URL          string
		Number       int
		Author       Author
		ChangedFiles int
		Additions    int
		Deletions    int
		Merged       bool
		Closed       bool
		CreatedAt    githubv4.DateTime
		UpdatedAt    githubv4.DateTime
		Commits      struct {
			Nodes []PullRequestCommit
		} `graphql:"commits(last:1)"`
		Reviews struct {
			Nodes []Review
		} `graphql:"reviews(last:5)"`
		Comments struct {
			Nodes []Comment
		} `graphql:"comments(last:5)"`
	}

	// Repository contains informationa about a GitHub repository
	Repository struct {
		Owner        Author
		Name         string
		PullRequests struct {
			Nodes []PullRequest
		} `graphql:"pullRequests(first: 10, states:OPEN, orderBy: {field: CREATED_AT, direction: DESC})"`
	}

	// Review contains information about a code review
	Review struct {
		State     string
		CreatedAt githubv4.DateTime
		Comments  struct {
			TotalCount int
		}
		Author Author
	}
)

type (
	// RepositoryData is a predefined query for getting repository data
	RepositoryData struct {
		Repository `graphql:"repository(owner:$owner, name:$name)"`
	}
)
