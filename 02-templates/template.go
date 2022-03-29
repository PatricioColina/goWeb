package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Usuarios struct {
	UserName string
	Edad     int
	Activo   bool
	Admin    bool
}

func Index(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(rw, " hola mundo")
	template, err := template.ParseFiles("index.html")
	usuario := Usuarios{"patricio", 36, true, false}

	if err != nil {
		panic(err)
	} else {
		//template.Execute(rw, nil)
		template.Execute(rw, usuario)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)

	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	fmt.Println("El servidot esta conectado al puerto 3000")
	fmt.Println("server: http://localhost:3000/")
	log.Fatal(server.ListenAndServe())

}
