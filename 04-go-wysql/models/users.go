package models

import "gomysql/db"

type User struct {
	Id       int64
	UserName string
	Password string
	Email    string
}

type Users []User

const UserSchame string = `CREATE TABLE users (
	id int(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	usrname VARCHAR(30) NOT NULL,
	password VARCHAR(100) NOT NULL,
	email VARCHAR(50),
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP )`

//construir usuario
func NewUsers(username, password, email string) *User {
	user := &User{UserName: username, Password: password, Email: email}

	return user
}

//crear usuario e insertar db
func CreateUsers(username, password, email string) *User {
	user := NewUsers(username, password, email)
	user.insert()
	return user
}

//inserta registro
func (user *User) insert() {
	sql := "INSERT users SET usrname=?,password=?,email=?"
	result, _ := db.Exec(sql, user.UserName, user.Password, user.Email)

	user.Id, _ = result.LastInsertId()

}

//listar registros
func ListUsers() Users {
	sql := "SELECT id, usrname,password, email FROM users"
	users := Users{}

	rows, _ := db.Query(sql)

	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.UserName, &user.Password, &user.Email)
		users = append(users, user)
	}

	return users
}
