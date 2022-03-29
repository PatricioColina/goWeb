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

func Saludar(nombre string) string {
	return "Hola " + nombre + " desde una funcion"
}

var templates = template.Must(template.New("T").ParseGlob("templates/**/*.html")) //de tener mas paginas se agregan aca

func renderTemplate(rw http.ResponseWriter, name string, data interface{}) {
	err := templates.ExecuteTemplate(rw, name, data)

	if err != nil {
		panic(err)
	}
}

func Index(rw http.ResponseWriter, r *http.Request) {

	c1 := Curso{"GO"}
	c2 := Curso{"Python"}
	c3 := Curso{"HTML"}
	c4 := Curso{"JAVA"}
	c5 := Curso{"C#"}

	//fmt.Fprintln(rw, " hola mundo")

	//template, err := template.ParseFiles("index.html", "base.html")
	//
	/*funciones := template.FuncMap{
		"saludar": Saludar,
	}*/
	//template, err := template.New("index.html").Funcs(funciones).ParseFiles("index.html")

	cursos := []Curso{c1, c2, c3, c4, c5}

	usuario := Usuarios{"patricio", 36, true, false, cursos}

	//template.Execute(rw, nil)
	//template.Execute(rw, usuario)
	/*err := templates.ExecuteTemplate(rw, "index.html", usuario)

	if err != nil {
		panic(err)
	}*/
	renderTemplate(rw, "index.html", usuario)

}

func Registro(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "registro.html", nil)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/registro", Registro)

	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	fmt.Println("El servidot esta conectado al puerto 3000")
	fmt.Println("server: http://localhost:3000/")
	log.Fatal(server.ListenAndServe())

}
