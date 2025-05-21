package git

import (
	"context"
	"encoding/base64"
	"encoding/json"
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
		log.BaseLogger.Error("Error fetching file content:", "err", err)
		return "", err
	}

	// Check if the path is a file
	if rawFileContent == nil || rawFileContent.GetType() != "file" {
		log.BaseLogger.Error("The provided path is not a file:", "path", path)
		return "", fmt.Errorf("the provided path is not a file: %s", path)
	}

	// Decode the file content (GitHub API returns it as base64-encoded)
	content, err := base64.StdEncoding.DecodeString(*rawFileContent.Content)
	if err != nil {
		log.BaseLogger.Error("Error decoding file content:", "err", err)
		return "", err
	}

	// Print the file content
	fileContent := string(content)
	log.BaseLogger.Debug("File Content", "content", fileContent)
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
		log.BaseLogger.Error("Error fetching directory content:", "err", err)
		return nil, err
	}

	// Check if the path is a directory
	if directoryContent == nil {
		log.BaseLogger.Error("The provided path is not a directory", "path", path)
		return nil, fmt.Errorf("the provided path is not a directory or is empty: %s", path)
	}

	// Collect file contents
	var fileContents []string
	for _, content := range directoryContent {
		if content.GetType() == "file" {
			// Fetch the file content
			fileContent, err := gh.ReadFileContent(owner, repo, branch, content.GetPath())
			if err != nil {
				log.BaseLogger.Error("Error fetching file content:", "file", content.GetPath(), "err", err)
				return nil, err
			}

			// Append the decoded content to the result
			fileContents = append(fileContents, fileContent)
		}
	}

	// Log and return the file contents
	log.BaseLogger.Debug("File contents in directory", "path", path, "fileContents", fileContents)
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
		log.BaseLogger.Error("File already exists", "filePath", filePath, "branch", branch)
		return err
	}
	if err != nil && !strings.Contains(err.Error(), "404") {
		log.BaseLogger.Error("Error checking file existence", "err", err)
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
		log.BaseLogger.Error("Error creating file:", "err", err)
		return err
	}

	log.BaseLogger.Debug("File created successfully", "fileName", filePath)
	return nil
}

func (gh *GitHub) CreateOrUpdateFile(owner string, repo string, branch string, filePath string, content string) error {
	// Check if the file exists
	fileContent, _, _, err := gh.Client.Repositories.GetContents(
		context.Background(),
		owner,
		repo,
		filePath,
		&github.RepositoryContentGetOptions{Ref: branch},
	)
	if err != nil && !strings.Contains(err.Error(), "404") {
		log.BaseLogger.Error("Error checking file existence:", "err", err)
		return err
	}

	// If the file exists, compare its content
	if fileContent != nil {
		decodedContent, decodeErr := fileContent.GetContent()
		if decodeErr != nil {
			log.BaseLogger.Error("Error decoding file content:", "err", decodeErr)
			return decodeErr
		}

		// Compare the existing content with the new content
		if decodedContent == content {
			log.BaseLogger.Info("File content is identical, no update needed", "fileName", filePath)
			return nil
		}

		// Prepare the file content options for update
		fileContentOptions := &github.RepositoryContentFileOptions{
			Message: github.String("Update " + filePath),
			Content: []byte(content),
			Branch:  github.String(branch),
			SHA:     github.String(fileContent.GetSHA()),
		}

		// Update the file in the repository
		_, _, err = gh.Client.Repositories.UpdateFile(
			context.Background(),
			owner,
			repo,
			filePath,
			fileContentOptions,
		)
		if err != nil {
			log.BaseLogger.Error("Error updating file:", "err", err)
			return err
		}

		log.BaseLogger.Debug("File updated successfully", "fileName", filePath)
		return nil
	}

	// If the file does not exist, create it
	fileContentOptions := &github.RepositoryContentFileOptions{
		Message: github.String("Create " + filePath),
		Content: []byte(content),
		Branch:  github.String(branch),
	}

	_, _, err = gh.Client.Repositories.CreateFile(
		context.Background(),
		owner,
		repo,
		filePath,
		fileContentOptions,
	)
	if err != nil {
		log.BaseLogger.Error("Error creating file:", "err", err)
		return err
	}

	log.BaseLogger.Debug("File created successfully", "fileName", filePath)
	return nil
}

func (gh *GitHub) TriggerPipeline(owner string, repo string, pipelineParams map[string]any) (string, error) {
	// Create a context
	ctx := context.Background()

	eventType := "Run Terraform"

	// Convert the client payload to JSON
	payloadBytes, err := json.Marshal(pipelineParams)
	if err != nil {
		log.BaseLogger.Error("Error marshalling client payload:", "err", err)
		return "", fmt.Errorf("failed to marshal client payload: %w", err)
	}

	// Create a repository dispatch request
	dispatchRequest := &github.DispatchRequestOptions{
		EventType:     eventType, // Custom event type
		ClientPayload: (*json.RawMessage)(&payloadBytes),
	}

	// Trigger the repository dispatch event
	_, res, err := gh.Client.Repositories.Dispatch(ctx, owner, repo, *dispatchRequest)
	if err != nil {
		log.BaseLogger.Error("Failed to trigger pipeline:", "err", err)
		return "", fmt.Errorf("failed to trigger pipeline: %w", err)
	}

	log.BaseLogger.Info("TriggerPipeline", "response", res)

	// Return the status of the dispatch
	return res.Status, nil
}

func (gh *GitHub) GetPipelineOutput(owner string, repo string, pipeline string) (string, error) {
	return "", nil
}
