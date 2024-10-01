## Pregunta 1

Has sido encargado de plantear una solución al siguiente problema.
Imagina que tenemos un sistema escrito en cualquier lenguaje de programación compilado,
este sistema se conecta a una base de datos SQL, el sistema funciona muy bien cuando hay
baja demanda de transacciones, pero cuando la cantidad de transacciones aumenta el sistema
deja de responder solicitudes, se encontró que la base de datos es el cuello de botella y no
acepta solicitudes en paralelo para completar las transacciones. ¿Qué solución plantearías
para receptar más cantidad de transacciones adaptándonos al cuello de botella de la base de
datos?

## Respuesta

```
a. Utilizaria el paquete "github.com/jackc/pgx/v5/pgxpool", que permite manejar un pool de conexiones del lado del servidor.
b. Tambien se podria utilizar el paquetes sync.Once que permite que solo se mantengan un hilo de conexion, pero solo lo usaria si tiene como fin una prueba de concepto (PoC).
```

## Pregunta 2

Has sido encargado de desarrollar una nueva forma de encriptar comunicaciones.
Básicamente, cada vocal del mensaje de entrada deberá ser precedida por otra cadena,
llamada la clave. La función recibirá dos parámetros de cadena: el primero será la clave y el
segundo, el mensaje. La función devolverá una cadena.

## Respuesta


```go
package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {

	//Dado que con un simple Scan no se puede leer texto con valores vacios, inicializare un scanner.
	scanner := bufio.NewScanner(os.Stdin)

	//CLAVE
	fmt.Println("Clave:")
	clave := ""

	//Leer el texto escaneado
	if scanner.Scan() {
		clave = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer la clave:", err)
	}

	//MENSAJE
	fmt.Println("Mensaje:")
	mensaje := ""

	//Leer el texto escaneado
	if scanner.Scan() {
		mensaje = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el mensaje:", err)
	}

	mensaje_codificado := Encrypt(clave, mensaje)

	fmt.Println(mensaje_codificado)

}
```

## Pregunta 3

Dado un arreglo de enteros, elimina todos los nodos consecutivos cuya suma sea cero y
devuelve los nodos restantes. Un arreglo vacío también puede ser un resultado válido. Si se
recibe un valor nulo, devuelve un arreglo vacío.

## Respuesta


```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RemoveConsecutives(array_inicial []int) []int {

	//Inicializamos la suma
	sum_actual := 0

	//Inicilizamos un flag
	posicion_a_eliminar := len(array_inicial) + 1

	//Mensaje a retornar
	nuevo_array := []int{}

	//Iteramos el array
	for i, numero := range array_inicial {

		//Si se tomo el valor de esta posicion, iterara al siguiente ciclo
		if i == posicion_a_eliminar {
			posicion_a_eliminar = len(array_inicial) + 1
			continue
		}

		//Sumamos el valor
		sum_actual += numero

		//Si la suma es 0
		if i != 0 && sum_actual == 0 {

			//Agregamos el valor de la siguiente posicion
			nuevo_array = append(nuevo_array, array_inicial[i+1])

			//Almacenamos la posicion del valor a ignorar (eliminar)
			if i+1 <= len(array_inicial) {
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
```

## Pregunta 4

Se te ha asignado poder entender esta arquitectura para poder aplicarlo a un
nuevo proyecto, se necesita que expliques a detalle cual es la función de cada
término mostrado en la imagen, en caso de desconocer de uno de los términos,
explicar lo que creas que significaría según el flujo que sigue cada proceso así
mismo si existiera cosas que mejorarías agrégalo como un detalle junto a su ¿Por
qué?

## Respuesta

```
[Explicacion]

1. La app web y mobile se conectan al API Gateway. El cual sirve como gestor de las solicitudes a los distintos microservicios, de igual forma, el API Gateway sirve como proxy inverso, por lo que no es necesario exponer la DNS/IP de los microservicios.

2. Cada microservicio tiene una BD dentro de su instancia de servidor.

3. Cada microservicio tiene una conexion con el broker, el cual sirve para enviar informacion entre microservicios a através de eventos (Event-Driven)

[Mejoras/Observaciones]

1. El API Gateway debe tener configurado el limite de llamadas por minuto, un api-key, actualizar los CORS.

2. Cada servicio debe tener una conexion HTTPS, un balanceador de carga, y el servidor debe conectarse a la BD con mediante un pool.

3. Cada servicio que es un consumer del broker debe tener forma de gestionar los errores, como devolver un Ack (false) y eliminar/bloquear la conexion a un canal automaticamente ante un error critico.
```

## Pregunta 5

Se te ha asignado la tarea de presentar una nueva tecnología que mejorará el
rendimiento del sistema de órdenes a un grupo de stakeholders que no son
técnicos. La tecnología implica el uso de microservicios para mejorar la
escalabilidad y reducir el tiempo de respuesta.

¿Cómo explicarías esta nueva tecnología y sus beneficios en términos no técnicos
para asegurar que todos comprendan su importancia y el impacto que tendrá en el
negocio?

## Respuesta

```
Como principal beneficio esta la reduccion de la "deuda técnica", es decir, en caso de desplegar el proyecto en un solo servidor, sera económico al tener solo un servidor y por ende su gestión será mas sencilla. Sin embargo, a medida que aumenten los clientes o se quieran agregar mas funcionalidades se van a necesitar mas recursos computacionales, es decir, aumentar la potencia del servidor, y para evitar un costo excesivo es que enves de tener un solo servidor con un exceso de capacidad que talvez algunas funciones solo necesiten pocas, es que se debe dividir en pequeños servidores. 

Ahora bien, al no haber planificado desde el inicio el tener muchos servidores, el codigo base del sistema fue construido para que funcione en un solo servidor, por lo que ahora se va a necesitar mas mano de obra para planificar la arquitectura de los servidores, del nuevo codigo y con un periodo de tiempo corto. En conclusión, si se inicia con solo un servidor, a mediano plazo cuando se quiera cambiar a microservicios, habra que sumar toda la inversión inicial que se hizo más todo el replanteo tecnico de implementar microservicios.
```
