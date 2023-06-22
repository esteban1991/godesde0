package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num int
var err error

func LeerPorTeclado2() {
	for i := 0; i < 1; {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Println("ingrese un  Número: ")

		if scanner.Scan() {
			num, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue

			} else {
				for x := 1; x < 11; x++ {
					fmt.Println(num * x)
				}
				i++
			}

		}
	}

}
