// Package muss main sein, weil es eine main-Funktion gibt.
package main

import "fmt"

// Funktion hat Return-Typ, aber kein return.
func PrintSomething(what string) {
	fmt.Print(what)
	fmt.Print("\n")
}

// numbers ist int, unten wird aber iteriert.
func ComputeProduct(numbers ...int) int {
	result := 1
	// (s.o.) Schleife mit int geht nicht.
	for _, num := range numbers {
		result *= num
	}
	return result
}

func main() {
	// Neue Variable nicht mit = definieren.
	p := ComputeProduct(1, 3, 5, 2, 0, 2)
	// Typ nicht beim Aufruf mit angeben.
	PrintSomething(fmt.Sprint(p))
}
