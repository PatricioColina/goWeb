package handlers

import (
	"apirest/db"
	"apirest/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(rw, "Lista todos los Usuarios")
	rw.Header().Set("Content-Type", "aplication/json")
	//rw.Header().Set("Content-Type", "text/xml")

	db.Connect()
	users := models.ListUsers()
	db.Close()
	//trasnformar a json
	//output, _ := json.Marshal(users)
	output, _ := json.Marshal(users)

	fmt.Fprintln(rw, string(output))

}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(rw, "obtiene un usuario")
	rw.Header().Set("Content-Type", "aplication/json")

	//obtener di
	vars := mux.Vars(r) //mapa tipo string
	userId, _ := strconv.Atoi(vars["id"])

	db.Connect()
	user := models.GetUser(userId)
	db.Close()

	output, _ := json.Marshal(user)

	fmt.Fprintln(rw, string(output))
}

func CreaUser(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(rw, "Crea un usuario")

	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		fmt.Fprintln(rw, http.StatusUnprocessableEntity)
	} else {
		db.Connect()
		user.Save()
		db.Close()
	}

	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))

}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(rw, "Actualiza usuario")
	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		fmt.Fprintln(rw, http.StatusUnprocessableEntity)
	} else {
		db.Connect()
		user.Save()
		db.Close()
	}

	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(rw, "Elimina Usuario")
	rw.Header().Set("Content-Type", "aplication/json")

	//obtener di
	vars := mux.Vars(r) //mapa tipo string
	userId, _ := strconv.Atoi(vars["id"])

	db.Connect()
	user := models.GetUser(userId)
	user.Delete()
	db.Close()

	output, _ := json.Marshal(user)

	fmt.Fprintln(rw, string(output))
}
