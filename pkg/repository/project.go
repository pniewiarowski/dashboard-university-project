package repository

import (
	"dashboard/pkg/database"
	"dashboard/pkg/model"
)

type ProjectRepository struct{}

func (_ *ProjectRepository) Create(project *model.Project) (*model.Project, error) {
	err := database.DataBase.Create(project).Error

	if err != nil {
		return nil, err
	}

	err = database.DataBase.First(&project).Error

	return project, err
}

func (_ *ProjectRepository) GetAll() ([]model.Project, error) {
	var projects []model.Project

	err := database.DataBase.Find(&projects).Error

	return projects, err
}

func (_ *ProjectRepository) GetByID(id uint) (model.Project, error) {
	var project model.Project

	err := database.DataBase.First(&project, "uuid = ?", id).Error

	return project, err
}

func (_ *ProjectRepository) Delete(project model.Project) {
	database.DataBase.Delete(&model.Project{}, project.UUID)
}
