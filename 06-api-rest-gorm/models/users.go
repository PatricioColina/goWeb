package models

import "gorm/db"

type User struct {
	Id       int64  `json:"id"` //alias para el jason
	UserName string `json:"userName"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Users []User

func MigrarUser() {
	db.Database.AutoMigrate(User{})
}
