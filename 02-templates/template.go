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
	Cursos   []Curso
}

type Curso struct {
	Nombre string
}

func Index(rw http.ResponseWriter, r *http.Request) {
	c1 := Curso{"GO"}
	c2 := Curso{"Python"}
	c3 := Curso{"HTML"}
	c4 := Curso{"JAVA"}
	c5 := Curso{"C#"}

	//fmt.Fprintln(rw, " hola mundo")
	template, err := template.ParseFiles("index.html")

	cursos := []Curso{c1, c2, c3, c4, c5}

	usuario := Usuarios{"patricio", 36, true, false, cursos}

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
