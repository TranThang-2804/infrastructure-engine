package bootstrap

import (
	"github.com/TranThang-2804/infrastructure-engine/internal/adapter/git"
)

type Application struct {
	GitStore git.GitStore
}

func App() Application {
	app := &Application{}
	app.GitStore = NewGitHubStore()
	return *app
}

func (app *Application) CloseDBConnection() {
	// CloseMongoDBConnection(app.Mongo)
}
