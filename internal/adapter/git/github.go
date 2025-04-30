package git

import (
	"context"
	"encoding/base64"

	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
	"github.com/google/go-github/v50/github"
)

type GitHub struct{}

func (gh *GitHub) ReadFileContent(path string, owner string, repo string, branch string) (string, error) {
	// Create a GitHub client
	client := github.NewClient(nil)

	// Fetch the file content
	rawFileContent, _, _, err := client.Repositories.GetContents(
		context.Background(),
		owner,
		repo,
		path,
		&github.RepositoryContentGetOptions{Ref: branch},
	)
	if err != nil {
		log.Logger.Error("Error fetching file content:", "err", err)
		return "", err
	}

	// Decode the file content (GitHub API returns it as base64-encoded)
	content, err := base64.StdEncoding.DecodeString(*rawFileContent.Content)
	if err != nil {
		log.Logger.Error("Error decoding file content:", "err", err)
		return "", err
	}

	// Print the file content
  fileContent := string(content)
  log.Logger.Info("File Content", "content", fileContent)
  return fileContent, nil
}
