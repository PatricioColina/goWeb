package handlers

import (
	"fmt"
	"net/http"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Lista todos los Usuarios")
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
