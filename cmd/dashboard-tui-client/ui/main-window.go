package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func GetMainWindow() *tview.Box {
	mainWindow := tview.NewBox()

	mainWindow.SetBorder(true)
	mainWindow.SetTitle("Dashboard Project")
	mainWindow.SetBorderColor(tcell.ColorDarkOrange)
	mainWindow.SetBackgroundColor(tcell.ColorSlateGray)
	mainWindow.SetTitleColor(tcell.ColorDarkOrange)

	return mainWindow
}
