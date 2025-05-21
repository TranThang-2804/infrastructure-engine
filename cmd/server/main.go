package main

import (
	"net/http"

	"github.com/TranThang-2804/infrastructure-engine/internal/api/route"
	"github.com/TranThang-2804/infrastructure-engine/internal/bootstrap"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/env"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
)

func init() {
	log.Init()
	env.LoadEnv()

	log.BaseLogger.Info("Starting the application - Author: Tommy Tran - tranthang.dev@gmail.com")
}

func main() {
	app := bootstrap.App()

	defer app.CloseDBConnection()

	r := route.SetupRoute(app)

	log.BaseLogger.Info("Starting server...", "on port", env.Env.ServerAddress)

	http.ListenAndServe(env.Env.ServerAddress, r)
}
