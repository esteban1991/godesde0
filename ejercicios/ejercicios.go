package ejercicios

import "strconv"

func ConviertoANum(palabra string) (int, string) {
	numero, err := strconv.Atoi(palabra)
	if err != nil {
		println(err)
		return numero, "error" + err.Error()
	} else {
		if numero > 100 {
			return numero, "es mayor a 100"
		} else {
			return numero, "es menor  a 100"
		}

	}

}
