package output1

import (
	"fmt"
	"strings"
)

/** AUFGABE:
  Geben Sie für jede der Funktionen in dieser Datei an,
  was sie auf der Konsole ausgeben.

  Ggf. sind bei den Funktionen noch Parameter angegeben,
  für die Sie die Ausgabe angeben sollen.
**/

// Aufgabe: Was wird ausgegeben?
// begin:Foo1
func Foo1() {
	x := 8
	y := 7
	k := x + 2*y - 2

	fmt.Println(x, y, k)
	fmt.Printf("%d-%d-%d\n", x, y, k)
	fmt.Println(k + 2 - x)
	fmt.Println(float64(k) / float64(x))
}

// end:Foo1

// Aufgabe: Was wird für n = 3, n = 4, n = 5 ausgegeben?
// begin:Foo2
func Foo2(n int) {
	s := "*"
	for i := 0; i < n; i++ {
		fmt.Println(s)
		s += s
	}
}

// end:Foo2

// Aufgabe: Was wird für n = 3, n = 4, n = 5 ausgegeben?
// begin:Foo3
func Foo3(n int) {
	if n == 0 {
		return
	}
	Foo3(n - 1)
	fmt.Println(strings.Repeat("*", n))
}

// end:Foo3

// Aufgabe: Was wird für n = 3, n = 4, n = 5 ausgegeben?
// begin:Foo4
func Foo4(n int) {
	fmt.Print(n)
	n = 42
	fmt.Println(n)
}

// end:Foo4

// Aufgabe: Was wird ausgegeben?
// begin:Foo5
func Foo5() {
	n := 23
	Foo4(n)
	fmt.Println(n)
}

// end:Foo5

// Aufgabe: Was wird ausgegeben?
// begin:Foo6
func Foo6() {
	x := 1
	{
		x := 2
		{
			x := 3
			fmt.Println(x)
		}
		fmt.Println(x)
	}
	fmt.Println(x)
}

// end:Foo6
