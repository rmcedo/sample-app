package main

import (
	"fmt"
)

var test = 03

//utilizado para package level scope
//tentar utilizar marmota sempre que der

// ----> VALORES ZERO <----
var a int
var b float64
var c string
var d bool

// Criação de Tipos de variaveis
type hotdog int

var e hotdog = 101

func main() {

	x := 10
	y := "Olá, bom dia"
	// Para uma raw string literals utilizamos o ' ' <--

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)

	// := apenas funciona para declarar uma variavél, não utilizada para atribuir um novo valor
	// := só pode ser utilizada dentro de um codeblock

	x, z := 20, 30
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n", z, z)

	qualquerCoisa(x)

	// ----> VALORES ZERO <----
	fmt.Printf("a: %v, %T\n", a, a)

	fmt.Printf("b: %v, %T\n", b, b)

	fmt.Printf("c: %v, %T\n", c, c)

	fmt.Printf("d: %v, %T\n", d, d)

	// ---> CRIAÇÃO DE TIPOS <---
	fmt.Printf("e: %v, %T\n", e, e)

	x = int(e)

	// ---> CONVERSÃO DE TIPOS <---
	fmt.Printf("x: %v, %T\n", x, x)

}

func qualquerCoisa(x int) {

	fmt.Println(test)
	fmt.Println(x)
}

/* EXERCÍCIO 3
	var x int = 42
var y string = "James Bond"
var z bool = true

func main() {

	s := fmt.Sprint(x, "\n", y, "\n", z)
	fmt.Println(s)

}
*/

/*EXERCÍCIO 4

type novoTipo int

var x novoTipo

func main() {

	fmt.Println(x)
	fmt.Printf("x: %T\n", x)
	x = 42
	fmt.Println(x)
	fmt.Printf("x: %T\n", x)

}
*/

/* EXERCÍCIO 5

type novoTipo int

var x novoTipo
var y int

func main() {

	fmt.Println(x)
	fmt.Printf("x: %T\n", x)
	x = 42
	fmt.Println(x)
	fmt.Printf("x: %T\n", x)

	y = int(x)

	fmt.Println(y)
	fmt.Printf("y: %T\n", y)

}

*/
