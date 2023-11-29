package page

import (
	"dashboard/pkg/model"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func BuildDashboardTextView(view *tview.TextView, user model.User) {
	view.SetBackgroundColor(tcell.ColorBlack)
	view.SetBorder(true)
	view.SetBorderColor(tcell.ColorDarkOrange)
	view.SetTitle("Dashboard - Main Menu")

	options := []string{
		fmt.Sprintf("-- Welcome %s %s! --", user.FirstName, user.LastName),
		"(s) show all your projects",
		"(a) add new project",
		"(i) check your account",
		"(l) logout",
	}

	render := ""
	for _, option := range options {
		render += fmt.Sprintf("%s\n", option)
	}

	view.SetText(render)
}
