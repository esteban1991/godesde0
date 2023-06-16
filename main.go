package main

import (
	"fmt"
	"github/godesde0/variables"
)

func main() {
	estado, texto := variables.ConviertoATexto(1588)
	fmt.Println(estado, texto)
}
