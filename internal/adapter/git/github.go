package git

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
	"github.com/google/go-github/v50/github"
)

type GitHub struct {
	Client *github.Client
}

func (gh *GitHub) ReadFileContent(owner string, repo string, branch string, path string) (string, error) {
	// Fetch the file content
	rawFileContent, _, _, err := gh.Client.Repositories.GetContents(
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

	// Check if the path is a file
	if rawFileContent == nil || rawFileContent.GetType() != "file" {
		log.Logger.Error("The provided path is not a file:", "path", path)
		return "", fmt.Errorf("the provided path is not a file: %s", path)
	}

	// Decode the file content (GitHub API returns it as base64-encoded)
	content, err := base64.StdEncoding.DecodeString(*rawFileContent.Content)
	if err != nil {
		log.Logger.Error("Error decoding file content:", "err", err)
		return "", err
	}

	// Print the file content
	fileContent := string(content)
	log.Logger.Debug("File Content", "content", fileContent)
	return fileContent, nil
}

func (gh *GitHub) GetAllFileContentsInDirectory(owner string, repo string, branch string, path string) ([]string, error) {
	// Fetch the directory content
	_, directoryContent, _, err := gh.Client.Repositories.GetContents(
		context.Background(),
		owner,
		repo,
		path,
		&github.RepositoryContentGetOptions{Ref: branch},
	)
	if err != nil {
		log.Logger.Error("Error fetching directory content:", "err", err)
		return nil, err
	}

	// Check if the path is a directory
	if directoryContent == nil {
		log.Logger.Error("The provided path is not a directory", "path", path)
		return nil, fmt.Errorf("the provided path is not a directory or is empty: %s", path)
	}

	// Collect file contents
	var fileContents []string
	for _, content := range directoryContent {
		if content.GetType() == "file" {
			// Fetch the file content
			fileContent, err := gh.ReadFileContent(owner, repo, branch, content.GetPath())
			if err != nil {
				log.Logger.Error("Error fetching file content:", "file", content.GetPath(), "err", err)
				return nil, err
			}

			// Append the decoded content to the result
			fileContents = append(fileContents, fileContent)
		}
	}

	// Log and return the file contents
	log.Logger.Debug("File contents in directory", "path", path, "fileContents", fileContents)
	return fileContents, nil
}

func (gh *GitHub) CreateFile(owner string, repo string, branch string, filePath string, content string) error {
	// Check if the file already exists
	fileContent, _, _, err := gh.Client.Repositories.GetContents(
		context.Background(),
		owner,
		repo,
		filePath,
		&github.RepositoryContentGetOptions{Ref: branch},
	)
	if err == nil && fileContent != nil {
		err := fmt.Errorf("file %s already exists in branch %s", filePath, branch)
		log.Logger.Error("File already exists", "filePath", filePath, "branch", branch)
		return err
	}
	if err != nil && !strings.Contains(err.Error(), "404") {
		log.Logger.Error("Error checking file existence", "err", err)
		return err
	}

	// Create the file content
	fileContentOptions := &github.RepositoryContentFileOptions{
		Message: github.String("Create " + filePath),
		Content: []byte(content),
		Branch:  github.String(branch),
	}

	// Create the file in the repository
	_, _, err = gh.Client.Repositories.CreateFile(
		context.Background(),
		owner,
		repo,
		filePath,
		fileContentOptions,
	)
	if err != nil {
		log.Logger.Error("Error creating file:", "err", err)
		return err
	}

	log.Logger.Debug("File created successfully", "fileName", filePath)
	return nil
}

func (gh *GitHub) CreateOrUpdateFile(owner string, repo string, branch string, filePath string, content string) error {
	// Create the file content
	fileContentOptions := &github.RepositoryContentFileOptions{
		Message: github.String("Create " + filePath),
		Content: []byte(content),
		Branch:  github.String(branch),
	}

	// Create the file in the repository
	_, _, err := gh.Client.Repositories.CreateFile(
		context.Background(),
		owner,
		repo,
		filePath,
		fileContentOptions,
	)
	if err != nil {
		log.Logger.Error("Error creating file:", "err", err)
		return err
	}

	log.Logger.Debug("File created successfully", "fileName", filePath)
	return nil
}
