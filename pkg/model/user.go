package model

type User struct {
	UUID      uint   `json:"uuid"`
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	Telephone string `json:"telephone"`
	Email     string `json:"email" gorm:"primaryKey"`
	Password  string `json:"password"`
	Enable    bool   `json:"enable"`
}
