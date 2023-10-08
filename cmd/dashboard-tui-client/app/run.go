package app

import (
	"dashboard/cmd/dashboard-tui-client/ui"
	"github.com/gofiber/fiber/v2/log"
	"github.com/rivo/tview"
)

func Run(restUrl string) {
	app := tview.NewApplication()

	app.SetRoot(ui.GetMainWindow(), true)

	log.Fatal(app.Run())
}
