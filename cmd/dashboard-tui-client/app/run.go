package app

import (
	"dashboard/cmd/dashboard-tui-client/page"
	"dashboard/pkg/ascii"
	"github.com/gdamore/tcell/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/rivo/tview"
)

const (
	MENU_PAGE = "menu_page"
	FORM_PAGE = "form_page"
)

func Run(restUrl string) {
	app := tview.NewApplication()
	pages := tview.NewPages()

	textView := tview.NewTextView()
	formView := tview.NewForm()

	pages.AddPage(MENU_PAGE, textView, true, true)
	pages.AddPage(FORM_PAGE, formView, true, true)

	app.SetRoot(pages, true)
	app.EnableMouse(true)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		current, _ := pages.GetFrontPage()

		if current == MENU_PAGE {
			if event.Rune() == ascii.KEY_ESCAPE {
				app.Stop()
			}

			if event.Rune() == ascii.KEY_LOWER_L {
				formView.Clear(true)
				page.BuildLoginForm(formView)

				pages.SwitchToPage(FORM_PAGE)
			}

			if event.Rune() == ascii.KEY_LOWER_R {
				formView.Clear(true)
				page.BuildRegisterForm(formView)

				pages.SwitchToPage(FORM_PAGE)
			}

			return event
		}

		if current == FORM_PAGE {
			if event.Rune() == ascii.KEY_ESCAPE {
				textView.Clear()
				page.BuildMainMenuTextView(textView)

				pages.SwitchToPage(MENU_PAGE)
			}

			return event
		}

		return event
	})

	page.BuildMainMenuTextView(textView)
	pages.SwitchToPage(MENU_PAGE)

	log.Fatal(app.Run())
}
