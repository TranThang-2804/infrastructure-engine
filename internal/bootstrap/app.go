package bootstrap

type Application struct {
	Env *Env
	// Mongo mongo.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	return *app
}

func (app *Application) CloseDBConnection() {
	// CloseMongoDBConnection(app.Mongo)
}
