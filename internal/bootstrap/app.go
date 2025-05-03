package bootstrap

import "github.com/TranThang-2804/infrastructure-engine/internal/adapter/git"

type Application struct {
	Env      *EnvConfig
	GitStore git.GitStore
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.GitStore = NewGitHubStore(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	// CloseMongoDBConnection(app.Mongo)
}
