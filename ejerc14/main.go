package main

import "log"

func main() {
	/*archivo := "prueba.txt"
	f, err := os.Open(archivo)

	defer f.Close() //esto hace que no se ejecute, solo cuando salga de la funcion
	if err != nil {
		fmt.Print("error al abrir fichero")
		os.Exit(1) //aun diciend que finalice el programa, el defer se ejecutara
	}*/

	ejemploPanic()
}

func ejemploPanic() {
	defer func() { //le añadimos defer al panic para poder jugar despues de panikear
		reco := recover()
		if reco != nil { //si reco es distinto a nil se ha dado un panic en el codigo
			log.Fatalf("ocurrio un error que genero panic", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("se encontro el valor de 1") //al hacer el panic se para el programa
	}
}
