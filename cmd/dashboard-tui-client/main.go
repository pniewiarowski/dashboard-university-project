package main

import (
	"dashboard/cmd/dashboard-tui-client/app"
	"dashboard/pkg/database"
	"dashboard/pkg/env"
	"dashboard/pkg/model"
)

var Models = []interface{}{
	model.User{},
	model.Project{},
}

func main() {
	env.Load(".env")

	database.Setup("dashboard")

	if env.GetMigration() {
		database.MakeMigration(Models)
	}

	app.Run()
}
