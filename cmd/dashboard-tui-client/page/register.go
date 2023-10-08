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

	form.AddInputField("First Name", "", 50, nil, func(firstName string) {
		user.FirstName = firstName
	})

	form.AddInputField("Last Name", "", 50, nil, func(lastName string) {
		user.LastName = lastName
	})

	form.AddInputField("Telephone", "", 50, nil, func(telephone string) {
		user.Telephone = telephone
	})

	form.AddInputField("Email", "", 50, nil, func(email string) {
		user.Email = email
	})

	form.AddPasswordField("Password", "", 50, 0, func(password string) {
		user.Password = password
	})

	form.AddCheckbox("Enable", false, func(enabled bool) {
		user.Enable = enabled
	})

	form.AddButton("Submit", func() {

	})
}
