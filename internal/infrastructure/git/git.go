package git

import "golang.org/x/net/context"

type GitStore interface {
	ReadFileContent(ctx context.Context, owner string, repo string, branch string, path string) (string, error)
	GetAllFileContentsInDirectory(
		ctx context.Context,
		owner string,
		repo string,
		branch string,
		path string,
	) ([]string, error)
	CreateFile(ctx context.Context, owner string, repo string, branch string, filePath string, content string) error
	CreateOrUpdateFile(
		ctx context.Context,
		owner string,
		repo string,
		branch string,
		filePath string,
		content string,
	) error
	TriggerPipeline(ctx context.Context, owner string, repo string, pipelineParams map[string]any) (string, error)
	GetPipelineOutput(ctx context.Context, owner string, repo string, pipeline string) (string, error)
}
