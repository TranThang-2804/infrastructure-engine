package git

type GitStore interface {
  ReadFileContent(path string, owner string, repo string, branch string) (string, error)
}
