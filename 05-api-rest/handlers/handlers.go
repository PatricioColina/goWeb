package handlers

import (
	"apirest/db"
	"apirest/models"
	"encoding/xml"
	"fmt"
	"net/http"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(rw, "Lista todos los Usuarios")
	//rw.Header().Set("Content-Type", "aplication/json")
	rw.Header().Set("Content-Type", "text/xml")

	db.Connect()
	users := models.ListUsers()
	db.Close()
	//trasnformar a json
	//output, _ := json.Marshal(users)
	output, _ := xml.Marshal(users)

	fmt.Fprintln(rw, string(output))

}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "obtiene un usuario")
}

func CreaUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Crea un usuario")
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Actualiza usuario")
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Elimina Usuario")
}
