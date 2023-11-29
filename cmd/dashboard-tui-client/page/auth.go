package page

import (
	"dashboard/pkg/model"
	"dashboard/pkg/repository"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func BuildUserView(view *tview.TextView, user model.User) {
	view.SetBackgroundColor(tcell.ColorBlack)
	view.SetBorder(true)
	view.SetBorderColor(tcell.ColorDarkOrange)
	view.SetTitle(user.FirstName + " " + user.LastName)

	render := ""
	render += fmt.Sprintf("User UUID:\t%d\n", user.UUID)
	render += fmt.Sprintf("User First Name:\t%s\n", user.FirstName)
	render += fmt.Sprintf("User Last Name:\t%s\n", user.LastName)
	render += fmt.Sprintf("User Telephone:\t%s\n", user.Telephone)
	render += fmt.Sprintf("User Email:\t%s\n", user.Email)
	render += "(esc) to go back"

	view.SetText(render)
}

func BuildAuthMenu(view *tview.TextView) {
	view.SetBackgroundColor(tcell.ColorBlack)
	view.SetBorder(true)
	view.SetBorderColor(tcell.ColorDarkOrange)
	view.SetTitle("Dashboard - Menu")

	options := []string{
		"(l) login to your account",
		"(r) or create account if you do not have one!",
	}

	render := ""
	for _, option := range options {
		render += fmt.Sprintf("%s\n", option)
	}

	view.SetText(render)
}

func BuildLoginForm(form *tview.Form, submit func(user model.User), cancel func()) {
	user := model.User{}

	form.SetBackgroundColor(tcell.ColorBlack)
	form.SetBorder(true)
	form.SetBorderColor(tcell.ColorDarkOrange)
	form.SetTitle("Dashboard - Login Form")

	form.AddInputField("User Email", "", 50, nil, func(email string) {
		user.Email = email
	})

	form.AddPasswordField("User Password", "", 50, 0, func(password string) {
		user.Password = password
	})

	form.AddButton("Login", func() {
		userRepository := repository.UserRepository{}
		verify, _ := userRepository.GetByEmail(user.Email)

		if verify.Email == user.Email && verify.Password == user.Password {
			submit(verify)
		}
	})

	form.AddButton("Cancel", cancel)
}

func BuildRegisterForm(form *tview.Form, submit func(), cancel func()) {
	user := model.User{}
	userRepository := repository.UserRepository{}

	form.SetBackgroundColor(tcell.ColorBlack)
	form.SetBorder(true)
	form.SetBorderColor(tcell.ColorDarkOrange)
	form.SetTitle("Dashboard - Register Your Account")

	form.AddInputField("User First Name", "", 50, nil, func(firstName string) {
		user.FirstName = firstName
	})

	form.AddInputField("User Last Name", "", 50, nil, func(lastName string) {
		user.LastName = lastName
	})

	form.AddInputField("User Telephone", "", 50, nil, func(telephone string) {
		user.Telephone = telephone
	})

	form.AddInputField("User Email", "", 50, nil, func(email string) {
		user.Email = email
	})

	form.AddPasswordField("User Password", "", 50, 0, func(password string) {
		user.Password = password
	})

	form.AddCheckbox("User Enable", false, func(enabled bool) {
		user.Enable = enabled
	})

	form.AddButton("Register", func() {
		userRepository.Create(&user)
		submit()
	})

	form.AddButton("Cancel", cancel)
}
