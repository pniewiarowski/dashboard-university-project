package main

import (
	"dashboard/cmd/dashboard-rest-server/app"
	"dashboard/pkg/database"
	"dashboard/pkg/env"
	"dashboard/pkg/model"
)

var Models = []interface{}{
	model.User{},
}

func main() {
	env.Load(".env")

	db := env.GetDb()
	migration := env.GetMigration()
	port := env.GetPort()

	database.Setup(db)
	if migration {
		database.MakeMigration(Models)
	}

	app.Run(int(port))
}
