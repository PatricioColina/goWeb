package models

type User struct {
	Id       int
	UserName string
	Password string
	Email    string
}

const UserSchame string = `CREATE TABLE users (
	id int(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	usrname VARCHAR(30) NOT NULL,
	password VARCHAR(100) NOT NULL,
	email VARCHAR(50),
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP )`
