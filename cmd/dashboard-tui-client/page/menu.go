package page

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func BuildMainMenuTextView(view *tview.TextView) {
	view.SetBackgroundColor(tcell.ColorBlack)
	view.SetBorder(true)
	view.SetBorderColor(tcell.ColorDarkOrange)
	view.SetTitle("Dashboard Project")

	options := []string{
		"(l) login to your account",
		"(r) or create account if you do not have one!",
		"(ESC) exit from application",
	}

	render := ""
	for _, option := range options {
		render += fmt.Sprintf("%s\n", option)
	}

	view.SetText(render)
}
