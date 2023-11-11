package output2

import "fmt"

/** AUFGABE:
  Geben Sie für die Funktion Foo() in dieser Datei an,
  was sie auf der Konsole ausgibt.

  Die Funktionen Bar() und Baz() werden von Foo() aufgerufen,
  müssen hier also nicht direkt beschrieben werden.
**/

// Aufgabe: Was wird ausgegeben?
// begin:task
func Foo() {
	x := 3
	y := 4
	fmt.Println(x, y)
	x = Bar(x, y)
	fmt.Println(x, y)
	Baz(x, y)
	fmt.Println(x, y)
}

// Bar ist eine Hilfsfunktion für Foo.
func Bar(x, y int) int {
	x = y - x
	return x
}

// Baz ist eine Hilfsfunktion für Foo.
func Baz(x, y int) {
	x, y = y, x
	fmt.Println(x, y)
}

// end:task
