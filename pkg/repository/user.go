package repository

import (
	"dashboard/pkg/database"
	"dashboard/pkg/model"
)

type UserRepository struct{}

func (_ *UserRepository) Create(project *model.User) (*model.User, error) {
	err := database.DataBase.Create(project).Error

	if err != nil {
		return nil, err
	}

	err = database.DataBase.First(&project).Error

	return project, err
}

func (_ *UserRepository) GetAll() ([]model.User, error) {
	var users []model.User

	err := database.DataBase.Find(&users).Error

	return users, err
}

func (_ *UserRepository) GetByEmail(email string) (model.User, error) {
	var user model.User

	err := database.DataBase.First(&user, "email = ?", email).Error

	return user, err
}
