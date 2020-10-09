package githubclient

import "github.com/shurcooL/githubv4"

type PageInfo struct {
	EndCursor   githubv4.String
	HasNextPage bool
}

type PullRequest struct {
	Title       githubv4.String
	BaseRefName githubv4.String
	Body        githubv4.String
}

type Milestone struct {
	ID           githubv4.String
	Number       githubv4.Int
	Title        githubv4.String
	Description  githubv4.String
	PullRequests struct {
		PageInfo PageInfo
		Nodes    []PullRequest
	} `graphql:"pullRequests(first: 1000, states: [MERGED])"`
}

type MilestonesQuery struct {
	Repository struct {
		Milestones struct {
			PageInfo PageInfo
			Nodes    []Milestone
		} `graphql:"milestones(first: 10, after: $milestoneCursor, orderBy: {field: NUMBER, direction: DESC}, states: [OPEN])"`
	} `graphql:"repository(owner: $owner, name: $name)"`
}
