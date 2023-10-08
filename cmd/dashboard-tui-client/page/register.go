package page

import (
	"dashboard/pkg/model"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func BuildRegisterForm(form *tview.Form) {
	user := model.User{}

	form.SetBackgroundColor(tcell.ColorBlack)
	form.SetBorder(true)
	form.SetBorderColor(tcell.ColorDarkOrange)
	form.SetTitle("Register Your Account")

	form.AddInputField("First Name", "", 20, nil, func(firstName string) {
		user.FirstName = firstName
	})

	form.AddInputField("Last Name", "", 20, nil, func(lastName string) {
		user.LastName = lastName
	})
}
