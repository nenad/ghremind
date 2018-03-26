package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	githubAPI       = "https://api.github.com"
	githubURL       = "https://github.com/"
	pullRequestPath = "/repos/%s/%s/pulls"
)

type (
	// Client is the HTTP client to access GitHub
	Client struct {
		client *http.Client
		token  string
	}
	// Author is information about the author of the PR
	Author struct {
		Username  string `json:"login"`
		AvatarURI string `json:"avatar_url"`
	}

	// Label is keeping the name of the label
	Label struct {
		Name  string `json:"name"`
		Color string `json:"color"`
	}

	// PullRequest keeps necessary data for a GitHub PR
	PullRequest struct {
		Number       int       `json:"number"`
		Title        string    `json:"title"`
		URL          string    `json:"html_url"`
		Labels       []Label   `json:"labels"`
		CreatedAt    time.Time `json:"created_at"`
		LastActivity time.Time `json:"updated_at"`
		Author       Author    `json:"user"`
	}

	// Repository keeps repository with PRs
	Repository struct {
		Name         string        `json:"name"`
		Owner        string        `json:"owner"`
		URL          string        `json:"url"`
		PullRequests []PullRequest `json:"pull_requests"`
	}
)

// NewClient provides a new GitHub client
func NewClient(token string) *Client {
	return &Client{
		http.DefaultClient,
		token,
	}
}

// GetRepositoryData returns all opened pull requests
// for a given owner and repository
func (c *Client) GetRepositoryData(owner, repository string) (*Repository, error) {
	path := fmt.Sprintf(pullRequestPath, owner, repository)
	url := fmt.Sprintf("%s%s", githubAPI, path)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if c.token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("token %s", c.token))
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var pullRequests []PullRequest
	err = json.NewDecoder(resp.Body).Decode(&pullRequests)

	if err != nil {
		return nil, err
	}

	repo := &Repository{
		URL:          githubURL + owner + "/" + repository,
		Name:         repository,
		Owner:        owner,
		PullRequests: pullRequests,
	}
	return repo, nil
}
