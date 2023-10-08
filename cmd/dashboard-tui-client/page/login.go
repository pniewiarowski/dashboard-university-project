package page

import (
	"dashboard/pkg/model"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func BuildLoginForm(form *tview.Form) {
	user := model.User{}

	form.SetBackgroundColor(tcell.ColorBlack)
	form.SetBorder(true)
	form.SetBorderColor(tcell.ColorDarkOrange)
	form.SetTitle("Register Your Account")

	form.AddInputField("Email", "", 50, nil, func(email string) {
		user.Email = email
	})

	form.AddPasswordField("Password", "", 50, 0, func(password string) {
		user.Password = password
	})

	form.AddButton("Login", func() {

	})
}
