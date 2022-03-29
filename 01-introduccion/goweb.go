//creacion de servidor
package main

import (
	"fmt"
	"log"
	"net/http"
)

//handler
func Hola(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("El metodo es: + ", r.Method)
	fmt.Fprintln(rw, " Hola mundo")
}

func pagina(rw http.ResponseWriter, r *http.Request) {
	http.NotFound(rw, r)
}

//se recomienda trabajar con los errores del usuario
func Error(rw http.ResponseWriter, r *http.Request) {
	//http.Error(rw,"La pagina no funciona",404)
	//http.Error(rw, "La pagina no funciona", http.StatusNotFound)
	http.Error(rw, "Este es un error ", http.StatusConflict)
}

func Saludar(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.URL.RawQuery)

	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	fmt.Fprintf(rw, "Hola, %s tu eded es %s", name, age)
	//http.Error(rw, "Este es un error ", http.StatusConflict)
}

func main() {
	//creacion de mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Hola)
	mux.HandleFunc("/page", pagina)
	mux.HandleFunc("/error", Error)
	mux.HandleFunc("/saludar", Saludar)

	//creacion de ruta router
	//http.HandleFunc("/", Hola)
	//http.HandleFunc("/page", pagina)
	//http.HandleFunc("/error", Error)
	//http.HandleFunc("/saludar", Saludar)

	//creacion un servidor
	//http.ListenAndServe("localhost:3000", nil)
	//fatal error
	//fmt.Println("El servidot esta conectado al puerto 3000")
	//fmt.Println("server: http://localhost:3000/")
	//log.Fatal(http.ListenAndServe("localhost:3000", nil))
	log.Fatal(http.ListenAndServe("localhost:3000", mux))

	//tambien se puede crear un servidor
	server := &http.Server{
		Addr:    "localhost/3000",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())

	//configurar el server para cada fez que se guarde o realice un cambio se hara un fresh
	//

}

/*
get para consguier
post guardar formularios y eso
put es para editar
delete es para eliminar
*/
