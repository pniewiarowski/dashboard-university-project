package main

import (
	"dashboard/cmd/dashboard-tui-client/app"
	"dashboard/pkg/env"
)

func main() {
	env.Load(".env")

	restUrl := env.GetRestUrl()

	app.Run(restUrl)
}
