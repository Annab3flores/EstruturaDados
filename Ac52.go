package main

import (
	"fmt"
)

func main() {

	var a, b, c float32
	var lados [3]float32
	fmt.Scan(&lados[0], &lados[1], &lados[2])

	a = lados[0]
	b = lados[1]
	c = lados[2]

	var aux float32
	if b > a {
		aux = a
		a = b
		b = aux
	}
	if c > a {
		aux = a
		a = c
		c = aux
	}
	if c > b {
		aux = b
		b = c
		c = aux
	}

	if a >= (b + c) {
		fmt.Println("NAO FORMA TRIANGULO")
	} else {
		if (a * a) == ((b * b) + (c * c)) {
			fmt.Println("TRIANGULO RETANGULO")
		}
		if (a * a) > ((b * b) + (c * c)) {
			fmt.Println("TRIANGULO OBTUSANGULO")
		}
		if (a * a) < ((b * b) + (c * c)) {
			fmt.Println("TRIANGULO ACUTANGULO")
		}
		if a == b && b == c {
			fmt.Println("TRIANGULO EQUILATERO")
		} else if a == b || b == c || a == c {
			fmt.Println("TRIANGULO ISOSCELES")
		}

	}
}
