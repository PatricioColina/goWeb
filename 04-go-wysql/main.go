package main

import (
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Connect()
	db.CreateTable(models.UserSchame, "users")
	db.Ping()
	db.Close()
	//db.Ping()//validacion erronea ya que la bd esta cerrada
}
