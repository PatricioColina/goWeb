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
	user := models.CreateUsers("alex", "pato123", "al@gmil.com")
	fmt.Println(user)
	db.Close()
	//db.Ping()//validacion erronea ya que la bd esta cerrada
}
