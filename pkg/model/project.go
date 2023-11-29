package model

type Project struct {
	UUID        uint   `json:"uuid"`
	Title       string `json:"title"`
	Enable      bool   `json:"enable"`
	Status      string `json:"status"`
	Description string `json:"description"`
}
