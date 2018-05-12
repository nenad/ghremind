package github

import (
	"context"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

// Client is GitHub GraphQL client with predefined queries
type Client struct {
	graphClient *githubv4.Client
	context     context.Context
}

// New returns new GitHub GraphQL client
func New(ctx context.Context, token string) *Client {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)
	return &Client{graphClient: client, context: ctx}
}

// RepositoryData returns data about repository
func (c *Client) RepositoryData(owner, name string) RepositoryData {
	var data RepositoryData
	vars := map[string]interface{}{
		"name":  githubv4.String(name),
		"owner": githubv4.String(owner),
	}

	c.graphClient.Query(c.context, &data, vars)
	return data
}
