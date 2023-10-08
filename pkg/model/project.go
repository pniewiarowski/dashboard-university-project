package model

import "dashboard/pkg/database"

type Project struct {
	EntityID uint   `json:"entity-id" gorm:"primaryKey"`
	Title    string `json:"title"`
	OwnerID  uint   `json:"owner-id"`
}

func GetProjectByEntityID(entityID uint) (Project, error) {
	var project Project

	err := database.DataBase.First(&project, entityID).Error

	return project, err
}

func GetAllProjects() ([]Project, error) {
	var projects []Project

	err := database.DataBase.Find(&projects).Error

	return projects, err
}

func CreateProject(project *Project) (*Project, error) {
	database.DataBase.Create(&project)
	err := database.DataBase.Find(*project).Error

	return project, err
}

func UpdateProject(project *Project, updateProject *Project) (*Project, error) {
	database.DataBase.Model(project).Updates(&updateProject)
	updateProject.EntityID = project.EntityID
	err := database.DataBase.Find(&updateProject).Error

	return updateProject, err
}

func DeleteProject(entityID uint) ([]Project, error) {
	var project Project
	var projects []Project

	err := database.DataBase.Delete(&project, entityID).Error
	if err != nil {
		database.DataBase.Find(&project)

		return []Project{}, err
	}
	err = database.DataBase.Find(&projects).Error

	return projects, err
}
