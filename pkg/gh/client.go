package gh

import (
	"context"
	"fmt"

	"github.com/google/go-github/v45/github"
)

type GitHubClientImpl interface {
	IsArchived(owner string, repo string) bool
	IsArchivedFromURL(string) bool
}

type repository struct {
	owner string
	name  string
}

type GitHubClient struct {
	client github.Client
}

func NewGitHubClient() *GitHubClient {
	return &GitHubClient{client: *github.NewClient(nil)}
}

func (gc *GitHubClient) IsArchived(owner string, repo string) (*bool, error) {
	ctx := context.Background()
	rep, _, err := gc.client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		return nil, fmt.Errorf("failed to get repository %w", err)
	}
	return rep.Archived, nil
}

func (gc *GitHubClient) IsArchivedFromURL(url string) (*bool, error) {
	repo, err := parseGitHubURL(url)
	if err != nil {
		return nil, err
	}
	return gc.IsArchived(repo.owner, repo.name)
}
