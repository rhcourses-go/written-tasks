package types2

import (
	"github.com/rhcourses-go/written-tasks/tests"
)

// Aufgabe: Welche Typen haben die Variablen in dieser Funktion?
func Types2() {
	// begin:task
	x := Foo1()
	y := Foo2()
	z := Foo4(x)
	z = z / 2.3
	if y {
		x += Foo3(z)
	}
	x += 3
	// end:task
	// Hilfsfunktion, nicht für die Aufgabe relevant.
	tests.PrintVarTypes(x, y, z)
}

// Foo1 gibt einen int zurück.
func Foo1() int {
	return 42
}

// Foo2 gibt einen bool zurück.
func Foo2() bool {
	return true
}

// Foo3 erwartet einen float64 und gibt einen int zurück.
func Foo3(float64) int {
	return 3
}

// Foo4 erwartet einen int und gibt einen float64 zurück.
func Foo4(int) float64 {
	return 0
}
