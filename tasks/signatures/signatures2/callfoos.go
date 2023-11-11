package signatures2

import "fmt"

/** AUFGABE:
  Welche Signaturen haben die Funktionen Foo1, Foo2, ...?
**/

// CallFoos ruft die Funktionen Foo1, Foo2, ... auf,
// und gibt die Typen der Rückgabewerte aus.
func CallFoos() {
	// begin:task
	x1 := Foo1(42)
	Foo2(x1)
	s1 := []string{"Hallo"}
	s2 := []string{"Welt"}
	if Foo3(s1) {
		x1 = Foo4(Foo3(s2), 15)
	}
	e1, e2 := Foo5(s1[0])
	fmt.Printf("%d\n", x1+e1+e2)
	// end:task
}

func Foo1(int) int {
	return 0
}

func Foo2(int) {}

func Foo3([]string) bool {
	return false
}

func Foo4(bool, int) int {
	return 0
}

func Foo5(string) (int, int) {
	return 0, 0
}
