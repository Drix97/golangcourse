package main

import (
	"fmt"
	"strings"
	"time"
)

//las go routines son lo equivalente a las promesas en node
func main() {
	go miNombreLento("Joel") //poniendo go delante de la llamada lo hace asincrono
	fmt.Println("estoy aqui")
	var x string
	fmt.Scanln(&x) //el problema es que al terminar el runtime, go no espera a que acabe la asincrona

}

func miNombreLento(nombre string) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Print(letra)
	}

}
