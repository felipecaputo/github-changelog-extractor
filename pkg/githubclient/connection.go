// Package githubclient provides queries and mutations to github V4 API
package githubclient

import (
	"context"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type Client struct {
	client *githubv4.Client
}

func NewClient(githubToken string) *Client {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

	return &Client{
		client: client,
	}
}

func (c *Client) GetClosedMilestones(repoOwner, repoName string, lastMilestoneCursor string) (*MilestonesQuery, error) {
	query := &MilestonesQuery{}

	variables := map[string]interface{}{
		"owner":           githubv4.String(repoOwner),
		"name":            githubv4.String(repoName),
		"milestoneCursor": (*githubv4.String)(nil),
	}

	if lastMilestoneCursor != "" {
		variables["milestoneCursor"] = githubv4.String(lastMilestoneCursor)
	}

	err := c.client.Query(context.Background(), query, variables)

	return query, err
}
