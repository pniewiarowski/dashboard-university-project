package page

import (
	"dashboard/pkg/model"
	"dashboard/pkg/repository"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
	_ "github.com/google/uuid"
	"github.com/rivo/tview"
)

func BuildProjectForm(form *tview.Form, submit func(), cancel func()) {
	project := model.Project{}
	projectRepository := repository.ProjectRepository{}

	form.SetBackgroundColor(tcell.ColorBlack)
	form.SetBorder(true)
	form.SetBorderColor(tcell.ColorDarkOrange)
	form.SetTitle("Dashboard - Register Your Account")

	form.AddInputField("Project Name", "", 50, nil, func(firstName string) {
		project.Title = firstName
	})

	form.AddCheckbox("Is Project Enabled", false, func(enabled bool) {
		project.Enable = enabled
	})

	form.AddDropDown("Project Status", []string{"completed", "in progress", "not started"}, 0, func(status string, idx int) {
		project.Status = status
	})

	form.AddTextArea("Project Description", "", 50, 4, 1024, func(description string) {
		project.Description = description
	})

	form.AddButton("Add", func() {
		id, _ := uuid.NewUUID()
		project.UUID = uint(id.ID())
		projectRepository.Create(&project)
		submit()
	})

	form.AddButton("Cancel", cancel)
}

func BuildProjectList(view *tview.List, submit func(id uint)) {
	projectRepository := repository.ProjectRepository{}
	projects, _ := projectRepository.GetAll()

	for idx, item := range projects {
		view.AddItem(item.Title, item.Status, rune(49+idx), func() {
			submit(item.UUID)
		})
	}
}

func BuildProjectView(view *tview.TextView, project model.Project) {
	view.SetBackgroundColor(tcell.ColorBlack)
	view.SetBorder(true)
	view.SetBorderColor(tcell.ColorDarkOrange)
	view.SetTitle(project.Title)

	render := ""
	render += fmt.Sprintf("Project Title:\t%s\n", project.Title)
	if project.Enable {
		render += "Project is enabled!\n"
	} else {
		render += "Project is disabled!\n"
	}
	render += fmt.Sprintf("Project Status:\t%s\n", project.Status)
	render += fmt.Sprintf("Project Description:\t%s\n", project.Description)
	render += "(esc) to go back"

	view.SetText(render)
}
