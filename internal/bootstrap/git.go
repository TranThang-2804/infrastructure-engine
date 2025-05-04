package bootstrap

import (
	"context"

	"github.com/TranThang-2804/infrastructure-engine/internal/infrastructure/git"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/env"
	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

func NewGitHubStore() git.GitStore {
	// Create an authenticated GitHub client using a personal access token
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: env.Env.GitToken},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	// Create a GitHub client
	client := github.NewClient(tc)

	return &git.GitHub{
		Client: client,
	}
}
