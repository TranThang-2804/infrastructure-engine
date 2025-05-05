package git

type GitStore interface {
	ReadFileContent(owner string, repo string, branch string, path string) (string, error)
	GetAllFileContentsInDirectory(owner string, repo string, branch string, path string) ([]string, error)
	CreateFile(owner string, repo string, branch string, filePath string, content string) error
	CreateOrUpdateFile(owner string, repo string, branch string, filePath string, content string) error
	TriggerPipeline(owner string, repo string, pipelinePayload []byte) (string, error)
	GetPipelineOutput(owner string, repo string, pipeline string) (string, error)
}
