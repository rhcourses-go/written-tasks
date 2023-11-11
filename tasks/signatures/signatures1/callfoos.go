package signatures1

/** AUFGABE:
  Welche Signaturen haben die Funktionen Foo1, Foo2, ...?
**/

// CallFoos ruft die Funktionen Foo1, Foo2, ... auf,
// und gibt die Typen der Rückgabewerte aus.
func CallFoos() bool {
	// begin:task
	x1 := Foo1("Hallo", 15)
	x2 := Foo2(x1)
	x3 := Foo3(127, x2)
	if x2 {
		x3 = append(x3, Foo1("Welt", x1))
	}
	x1 += Foo4(x2, true)
	Foo5(x2 && x1 != Foo4(x2, x2))
	return x2 && !(x1 > x3[0])
	// end:task
}
