package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//generacion de conexion
//username:password@tcp(localhost:3306)/database
const url = "root;root@tcp(localhost:3306)/goWeb_db"

var db *sql.DB

func Connect() {
	conection, err := sql.Open("mysql", url)

	if err != nil {
		panic(err)
	}

	db = conection

}

func Close() {
	db.Close()
}
