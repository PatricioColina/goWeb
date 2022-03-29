package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//generacion de conexion
//username:password@tcp(localhost:3306)/database
const url = "root:root@tcp(localhost:3306)/goweb_db"

//db, err := sql.Open("mysql", "user:password@/dbname")

//guarda la conexion
var db *sql.DB

//realiza la conexion
func Connect() {
	conection, err := sql.Open("mysql", url)

	if err != nil {
		panic(err)
	}
	fmt.Println("conexion exitosa")
	db = conection

}

//cierra la conexion
func Close() {
	db.Close()
}

//verificar la conexion
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}
