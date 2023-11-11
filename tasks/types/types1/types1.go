package types1

import "github.com/rhcourses-go/written-tasks/tests"

// Aufgabe: Welche Typen haben die Variablen in dieser Funktion?
func Types1() {
	// begin:task
	x := 42
	y := x - 2
	z := 3.14
	w := "Hello"
	// end:task

	tests.PrintVarTypes(x, y, z, w)
}
