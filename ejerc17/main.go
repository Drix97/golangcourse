package main

import "net/http"

//servidor web http

func main() {
	http.HandleFunc("/", home) //cuando en mi pagina vaya a barra(osea home) que ejecute la funcion home
	//	http.HandleFunc("/login", login) s

	http.ListenAndServe(":3000", nil) //escucha el puerto 3000 hasta que se triggeree
}

func home(w http.ResponseWriter, r *http.Request) { //estos seran siempre los dos parametros para conectar
	http.ServeFile(w, r, "./index.html") //cuando se ejecute home, que muestre index.html
}
