package model

import (
	"dashboard/pkg/database"
)

type User struct {
	EntityID  uint   `json:"entity-id" gorm:"primaryKey"`
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	Telephone string `json:"telephone"`
	Email     string `json:"email" gorm:"primaryKey"`
	Password  string `json:"password"`
	Enable    bool   `json:"enable"`
}

func CreateUser(user *User) error {
	return database.DataBase.Create(&user).Error
}

func GetUserByEntityID(entityID uint) *User {
	user := new(User)
	database.DataBase.First(&user, "entity_id = ?", entityID)

	return user
}

func GetUserByEmail(email string) *User {
	user := new(User)
	database.DataBase.First(&user, "email = ?", email)

	return user
}
