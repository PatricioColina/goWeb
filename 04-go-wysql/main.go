package main

import (
	"gomysql/db"
)

func main() {
	db.Connect()
	//db.CreateTable(models.UserSchame, "users")
	//db.Ping()
	//db.TruncateTable("users")
	//user := models.CreateUsers("Patricio", "pato123", "al@gmil.com")
	//users := models.ListUsers()

	//fmt.Println(users)

	/*user := models.GetUser(2)
	fmt.Println(user)

	user.Delete()
	//user.UserName = "Juan"
	//user.Save()
	fmt.Println(user)
	users := models.ListUsers()
	fmt.Println(users)*/

	db.TruncateTable("users")
	db.Close()
	//db.Ping()//validacion erronea ya que la bd esta cerrada
}
