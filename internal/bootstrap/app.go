package bootstrap

import (
	"github.com/TranThang-2804/infrastructure-engine/internal/adapter/git"
)

type Application struct {
	GitStore      git.GitStore
	InfraPipeline InfraPipeline
}

func App() Application {
	app := &Application{}
	app.GitStore = NewGitHubStore()
	app.InfraPipeline = *NewInfraPipeline(app.GitStore)

	app.InfraPipeline.SettingInfraPipeline()
	return *app
}

func (app *Application) CloseDBConnection() {
	// CloseMongoDBConnection(app.Mongo)
}
