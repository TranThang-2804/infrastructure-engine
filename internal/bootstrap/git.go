package bootstrap

import "github.com/TranThang-2804/infrastructure-engine/internal/adapter/git"

func NewGitHubStore(env *Env) git.GitStore {
	return &git.GitHub{}
}
