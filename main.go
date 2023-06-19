package main

import (
	"fmt"
	"github/godesde0/ejercicios"
	"github/godesde0/variables"
	"runtime"
)

func main() {
	estado, texto := variables.ConviertoATexto(1588)
	fmt.Println(estado, texto)
	if os := runtime.GOOS; os == "linux" || os == "OS X." {
		fmt.Println("esto no es windows")

	} else {
		fmt.Println("esto es windows ")
	}

	num, mensaje := ejercicios.ConviertoANum("50")
	fmt.Println(num, mensaje)

}
