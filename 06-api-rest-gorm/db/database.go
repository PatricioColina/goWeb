package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//"user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
var dsn = "root:root@tcp(localhost:3306)/goweb_db?charset=utf8mb4&parseTime=True&loc=Local"

var Database = func() (db *gorm.DB) {
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Error en connexion ", err)
		panic(err)
	} else {
		fmt.Println("Conexion exitosa")
		return db
	}
}()
