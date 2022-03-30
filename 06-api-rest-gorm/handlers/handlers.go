package handlers

import (
	"encoding/json"
	"gorm/db"
	"gorm/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(rw, "Lista todos los Usuarios")
	users := models.User{}
	db.Database.Find(&users)
	sendData(rw, users, http.StatusOK)

}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(rw, "obtiene un usuario")
	if user, err := getUserById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		sendData(rw, user, http.StatusOK)
	}
}

func getUserById(r *http.Request) (models.User, *gorm.DB) {

	vars := mux.Vars(r) //mapa tipo string
	userId, _ := strconv.Atoi(vars["id"])
	user := models.User{}
	if err := db.Database.First(&user, userId); err.Error != nil {
		return user, err
	} else {
		return user, nil
	}

}

func CreaUser(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(rw, "Crea un usuario")
	user := models.User{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		sendError(rw, http.StatusUnprocessableEntity)
	} else {
		db.Database.Save(&user)
		sendData(rw, user, http.StatusOK)
	}

}

//actualiza usuario
func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(rw, "Actualiza usuario")
	var userId int64

	if user_ant, err := getUserById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		userId = user_ant.Id
	}

	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		sendError(rw, http.StatusUnprocessableEntity)
	} else {
		user.Id = userId
		db.Database.Save(&user)
		sendData(rw, user, http.StatusOK)
	}
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(rw, "Elimina Usuario")
	if user, err := getUserById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		db.Database.Delete(&user)
		sendData(rw, user, http.StatusOK)
	}
}
