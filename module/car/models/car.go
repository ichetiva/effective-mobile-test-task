package models

import "gorm.io/gorm"

type Car struct {
	ID      uint   `gorm:"primarykey"`
	RegNum  string `json:"regNum"`
	Mark    string `json:"mark"`
	Model   string `json:"model"`
	Year    string `json:"year"`
	OwnerID uint
	Owner   Owner `json:"owner" gorm:"foreignKey:OwnerID"`
}

type Owner struct {
	gorm.Model
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
