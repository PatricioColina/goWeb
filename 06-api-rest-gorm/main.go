package main

import (
	"gorm/models"
)

func main() {
	models.MigrarUser()
	/*
		//rutas
		mux := mux.NewRouter()

		//endPoint
		mux.HandleFunc("/api/user/", handlers.GetUsers).Methods("GET")
		mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUser).Methods("GET")

		mux.HandleFunc("/api/user/", handlers.CreaUser).Methods("POST")
		mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")    //editar
		mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE") //eliminar

		log.Fatal(http.ListenAndServe(":3000", mux))
	*/
}
