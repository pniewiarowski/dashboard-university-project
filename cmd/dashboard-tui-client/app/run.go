package app

import (
	"dashboard/cmd/dashboard-tui-client/page"
	"dashboard/pkg/ascii"
	"dashboard/pkg/model"
	"dashboard/pkg/repository"
	"github.com/gdamore/tcell/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/rivo/tview"
)

const (
	MENU_PAGE = "menu_page"
	FORM_PAGE = "form_page"
	LIST_PAGE = "list_page"
)

type State struct {
	IsLogged bool
	User     model.User
}

func Run() {
	app := tview.NewApplication()
	pages := tview.NewPages()

	state := State{IsLogged: false}

	textView := tview.NewTextView()
	formView := tview.NewForm()
	listView := tview.NewList()

	pages.AddPage(MENU_PAGE, textView, true, true)
	pages.AddPage(FORM_PAGE, formView, true, true)
	pages.AddPage(LIST_PAGE, listView, true, true)

	app.SetRoot(pages, true)
	app.EnableMouse(true)

	projectRepository := repository.ProjectRepository{}

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		current, _ := pages.GetFrontPage()

		if state.IsLogged {
			if current == MENU_PAGE {
				if event.Rune() == ascii.KEY_LOWER_S {
					listView.Clear()
					page.BuildProjectList(listView, func(id uint) {
						textView.Clear()
						project, _ := projectRepository.GetByID(id)
						page.BuildProjectView(textView, project)
						pages.SwitchToPage(MENU_PAGE)
					})
					pages.SwitchToPage(LIST_PAGE)
				} else if event.Rune() == ascii.KEY_ESCAPE {
					textView.Clear()
					page.BuildDashboardTextView(textView, state.User)

					pages.SwitchToPage(MENU_PAGE)
				} else if event.Rune() == ascii.KEY_LOWER_A {
					formView.Clear(true)
					page.BuildProjectForm(formView, func() {
						listView.Clear()
						page.BuildProjectList(listView, func(id uint) {
							textView.Clear()
							project, _ := projectRepository.GetByID(id)
							page.BuildProjectView(textView, project)
							pages.SwitchToPage(MENU_PAGE)
						})
						pages.SwitchToPage(LIST_PAGE)
					}, func() {
						textView.Clear()
						page.BuildDashboardTextView(textView, state.User)
						pages.SwitchToPage(MENU_PAGE)
					})

					pages.SwitchToPage(FORM_PAGE)
				} else if event.Rune() == ascii.KEY_LOWER_L {
					textView.Clear()
					page.BuildAuthMenu(textView)
					state.IsLogged = false

					pages.SwitchToPage(MENU_PAGE)
				} else if event.Rune() == ascii.KEY_LOWER_I {
					textView.Clear()
					page.BuildUserView(textView, state.User)
					pages.SwitchToPage(MENU_PAGE)
				}

				return event
			}

			if current == FORM_PAGE {
				if event.Rune() == ascii.KEY_ESCAPE {
					textView.Clear()
					page.BuildDashboardTextView(textView, state.User)

					pages.SwitchToPage(MENU_PAGE)
				}
			}

			if current == LIST_PAGE {
				if event.Rune() == ascii.KEY_ESCAPE {
					textView.Clear()
					page.BuildDashboardTextView(textView, state.User)

					pages.SwitchToPage(MENU_PAGE)
				}
			}

			return event
		}

		if !state.IsLogged {
			if current == MENU_PAGE {
				if event.Rune() == ascii.KEY_ESCAPE {
					app.Stop()
				} else if event.Rune() == ascii.KEY_LOWER_L {
					formView.Clear(true)
					page.BuildLoginForm(formView, func(user model.User) {
						textView.Clear()
						state.IsLogged = true
						state.User = user

						page.BuildDashboardTextView(textView, user)

						pages.SwitchToPage(MENU_PAGE)
					}, func() {
						pages.SwitchToPage(MENU_PAGE)
					})

					pages.SwitchToPage(FORM_PAGE)
				} else if event.Rune() == ascii.KEY_LOWER_R {
					formView.Clear(true)
					page.BuildRegisterForm(formView, func() {
						pages.SwitchToPage(MENU_PAGE)
					}, func() {
						pages.SwitchToPage(MENU_PAGE)
					})

					pages.SwitchToPage(FORM_PAGE)
				}
			} else if current == FORM_PAGE {
				if event.Rune() == ascii.KEY_ESCAPE {
					textView.Clear()
					page.BuildAuthMenu(textView)

					pages.SwitchToPage(MENU_PAGE)
				}
			}

			return event
		}

		return event
	})

	page.BuildAuthMenu(textView)
	pages.SwitchToPage(MENU_PAGE)

	log.Fatal(app.Run())
}
