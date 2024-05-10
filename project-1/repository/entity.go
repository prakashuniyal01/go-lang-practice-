package repository

import "gorm.io/gorm"

type UserData struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Age       int
	Married   bool
}
