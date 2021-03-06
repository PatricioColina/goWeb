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

//CREA UNA TABLA
func CreateTable(schema string, name string) {
	if !ExitsTable(name) {
		_, err := db.Exec(schema)

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

}

//reiniciar tabla
func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE TABLE %s", tableName)
	Exec(sql)
}

//funcion que valida la tabla
func ExitsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLE LIKE '%s'", tableName)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return rows.Next()
}

//polimorfismo de exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	Connect()
	resust, err := db.Exec(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}

	return resust, err
}

//polimorfismo de query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	Connect()
	rows, err := db.Query(query, args...)
	Close()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return rows, err
}
