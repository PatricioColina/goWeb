package handlers

import (
	"apirest/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(rw, "Lista todos los Usuarios")
	if users, err := models.ListUsers(); err != nil {
		models.SendNoFound(rw)
	} else {
		models.SendData(rw, users)
	}

}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(rw, "obtiene un usuario")
	if user, err := getUserByRequest(r); err != nil {
		models.SendNoFound(rw)
	} else {
		models.SendData(rw, user)
	}
}

func CreaUser(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(rw, "Crea un usuario")

	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		models.SendStatusUnprocessableEntity(rw)
	} else {
		user.Save()
		models.SendData(rw, user)
	}

}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(rw, "Actualiza usuario")
	var userId int64

	if user, err := getUserByRequest(r); err != nil {
		models.SendNoFound(rw)
	} else {
		userId = user.Id
	}

	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		models.SendStatusUnprocessableEntity(rw)
	} else {
		user.Id = userId
		user.Save()
		models.SendData(rw, user)
	}
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(rw, "Elimina Usuario")
	if user, err := getUserByRequest(r); err != nil {
		models.SendNoFound(rw)
	} else {
		user.Delete()
		models.SendData(rw, user)
	}
}

func getUserByRequest(r *http.Request) (models.User, error) {
	vars := mux.Vars(r) //mapa tipo string
	userId, _ := strconv.Atoi(vars["id"])

	if user, err := models.GetUser(userId); err != nil {
		return *user, err
	} else {
		return *user, nil
	}

}
