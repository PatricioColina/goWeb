package main

import (
	"gomysql/db"
)

func main() {
	db.Connect()
	//db.CreateTable(models.UserSchame, "users")
	//db.Ping()
	db.TruncateTable("users")
	db.Close()
	//db.Ping()//validacion erronea ya que la bd esta cerrada
}
