package main

import (
	"fmt"
	"time"
)

//un canal es un espacio en memoria de un tipo de dato
func main() {
	canal1 := make(chan time.Duration)
	go bucle(canal1)
	fmt.Println("Lklegue hasta aqui")

	msg := <-canal1 //aqui espera a que canal1 tenga un valor y lo grabe en msg. Hasta que no ocurre, el programa no continua
	fmt.Println(msg)
}

func bucle(canal1 chan time.Duration) {
	inicio := time.Now()
	for i := 0; i < 100000000000; i++ {

	}
	final := time.Now()
	canal1 <- final.Sub(inicio) //Sub duevuelve la duracion y es un tipo de dato
}
