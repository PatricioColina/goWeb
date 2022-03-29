package main

import (
	"fmt"
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Connect()
	//db.CreateTable(models.UserSchame, "users")
	//db.Ping()
	//db.TruncateTable("users")
	//user := models.CreateUsers("Patricio", "pato123", "al@gmil.com")
	users := models.ListUsers()
	fmt.Println(users)
	//fmt.Println(user)
	db.Close()
	//db.Ping()//validacion erronea ya que la bd esta cerrada
}
