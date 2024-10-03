package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Encrypt(clave string, mensaje string) string {

	if clave == "" {
		clave = "DCJ"
	}

	//Mensaje a retornar
	//El uso de strings.Builder es para evitar la creacion de multiples cadenas de texto debido a la inmutabilidad de estas
	var mensaje_codificado strings.Builder

	//Iteramos el array
	for _, rune_caracter := range mensaje {

		siguiente_valor := ""

		if strings.Contains("aAeEiIoOuU", string(rune_caracter)) {
			siguiente_valor = clave
		}

		//Concatenamos el valor
		mensaje_codificado.WriteString(siguiente_valor + string(rune_caracter))
	}

	return mensaje_codificado.String()
}

func RemoveConsecutives(array_inicial []int) []int {

	//Inicializamos las variables
	sum_actual := 0
	longitud_array_incial := len(array_inicial)

	//Inicilizamos un flag
	posicion_a_eliminar := longitud_array_incial + 1

	//Mensaje a retornar
	nuevo_array := []int{}

	//Iteramos el array
	for i, numero := range array_inicial {

		//Si se tomo el valor de esta posicion, iterara al siguiente ciclo
		if i == posicion_a_eliminar {
			posicion_a_eliminar = longitud_array_incial + 1
			continue
		}

		//Sumamos el valor
		sum_actual += numero

		//Si la suma es 0
		if i != 0 && sum_actual == 0 {

			//Agregamos el valor de la siguiente posicion
			nuevo_array = append(nuevo_array, array_inicial[i+1])

			//Almacenamos la posicion del valor a ignorar (eliminar)
			if i+1 <= longitud_array_incial {
				posicion_a_eliminar = i + 1
			}

		}

	}

	return nuevo_array
}

func main() {

	//Dado que con un simple Scan no se puede leer texto con valores vacios, inicializare un scanner.
	scanner := bufio.NewScanner(os.Stdin)

	//ARRAY INICIAL
	fmt.Println("Array inicial:")
	var array_numeros []int

	//Leer el texto escaneado
	if scanner.Scan() {
		array_inicial := scanner.Text()
		numerosStr := strings.Split(array_inicial, ",")

		// Convertir las palabras a números enteros
		for _, numeroStr := range numerosStr {
			numero, err := strconv.Atoi(numeroStr)
			if err != nil {
				fmt.Println("Error al convertir:", err)
				return
			}
			array_numeros = append(array_numeros, numero)
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer la clave:", err)
	}

	nuevo_array := RemoveConsecutives(array_numeros)

	fmt.Println(nuevo_array)
}
