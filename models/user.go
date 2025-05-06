package models

type User struct {
	Id        uint `gorm:"unique" `
	FirstName string
	LastName  string
	Email     string `gorm:"unique" `
	Password  string
}
