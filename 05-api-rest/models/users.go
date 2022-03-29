package models

import "apirest/db"

type User struct {
	/*
		Id       int64  `json:"id"` //alias para el jason
		UserName string `json:"userName"`
		Password string `json:"password"`
		Email    string `json:"email"`
	*/
	Id       int64  `xml:"id"` //alias para el jason
	UserName string `xml:"userName"`
	Password string `xml:"password"`
	Email    string `xml:"email"`
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

//obtener un registro
func GetUser(id int) *User {
	user := NewUsers("", "", "")
	sql := "SELECT id, usrname,password, email FROM users WHERE id=?"
	rows, _ := db.Query(sql, id)
	for rows.Next() {
		rows.Scan(&user.Id, &user.UserName, &user.Password, &user.Email)
	}

	return user
}

//editar registro
func (user *User) update() {
	sql := "UPDATE users SET usrname=?,password=?,email=? WHERE id=?"

	db.Exec(sql, user.UserName, user.Password, user.Email, user.Id)

}

//guardar o editar registro
func (user *User) Save() {
	if user.Id == 0 {
		user.insert()
	} else {
		user.update()
	}
}

//eliminar registro
func (user *User) Delete() {
	sql := "DELETE FROM users WHERE id=?"
	db.Exec(sql, user.Id)
}
