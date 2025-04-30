package git

type GitStore interface {
  ReadFileContent(owner string, repo string, branch string, path string) (string, error)
  GetAllFileContentsInDirectory(owner string, repo string, branch string, path string) ([]string, error)
}
